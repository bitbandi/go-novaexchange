package novaexchange

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
	"net/url"
)

type client struct {
	apiKey      string
	apiSecret   string
	httpClient  *http.Client
	httpTimeout time.Duration
	debug       bool
}

// NewClient return a new NovaExchange HTTP client
func NewClient(apiKey, apiSecret string) (c *client) {
	return &client{apiKey, apiSecret, &http.Client{}, 30 * time.Second, false}
}

// NewClientWithCustomHttpConfig returns a new NovaExchange HTTP client using the predefined http client
func NewClientWithCustomHttpConfig(apiKey, apiSecret string, httpClient *http.Client) (c *client) {
	timeout := httpClient.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	return &client{apiKey, apiSecret, httpClient, timeout, false}
}

// NewClient returns a new NovaExchange HTTP client with custom timeout
func NewClientWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) (c *client) {
	return &client{apiKey, apiSecret, &http.Client{}, timeout, false}
}

func (c client) dumpRequest(r *http.Request) {
	if r == nil {
		log.Print("dumpReq ok: <nil>")
		return
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("dumpReq err:", err)
	} else {
		log.Print("dumpReq ok:", string(dump))
	}
}

func (c client) dumpResponse(r *http.Response) {
	if r == nil {
		log.Print("dumpResponse ok: <nil>")
		return
	}
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Print("dumpResponse err:", err)
	} else {
		log.Print("dumpResponse ok:", string(dump))
	}
}

// doTimeoutRequest do a HTTP request with timeout
func (c *client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	// Do the request in the background so we can check the timeout
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		if c.debug {
			c.dumpRequest(req)
		}
		resp, err := c.httpClient.Do(req)
		if c.debug {
			c.dumpResponse(resp)
		}
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("timeout on reading data from NovaExchange API")
	}
}

// do prepare and process HTTP request to NovaExchange API
func (c *client) do(method string, ressource string, payload string, authNeeded bool) (response []byte, err error) {
	connectTimer := time.NewTimer(c.httpTimeout)

	var rawurl string
	if strings.HasPrefix(ressource, "http") {
		rawurl = ressource
	} else {
		rawurl = fmt.Sprintf("%s%s/%s", API_BASE, API_VERSION, ressource)
	}

	// Auth
	if authNeeded {
		if len(c.apiKey) == 0 || len(c.apiSecret) == 0 {
			err = errors.New("You need to set API Key and API Secret to call this method")
			return
		}
		nonce := time.Now().UnixNano()
		var URL *url.URL
		URL, err = url.Parse(rawurl)
		if err != nil {
			return
		}
		q := URL.Query()
		q.Set("nonce", fmt.Sprintf("%d", nonce))
		URL.RawQuery = q.Encode()
		rawurl = URL.String()
		mac := hmac.New(sha512.New, []byte(c.apiSecret))
		_, err = mac.Write([]byte(URL.String()))
		sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		var values url.Values
		values, err = url.ParseQuery(payload)
		if err != nil {
			return
		}
		values.Add("apikey", c.apiKey)
		values.Add("signature", sig)
		payload = values.Encode()
	}

	req, err := http.NewRequest(method, rawurl, strings.NewReader(payload))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.doTimeoutRequest(connectTimer, req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	//fmt.Println(fmt.Sprintf("reponse %s", response), err)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return response, err
}
