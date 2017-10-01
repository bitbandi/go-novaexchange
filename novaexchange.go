// Package Novaexchange is an implementation of the Novaexchange API in Golang.
package novaexchange

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
	"fmt"
	"log"
)

const (
	API_BASE    = "https://novaexchange.com/remote/" // NovaExchange API endpoint
	API_VERSION = "v2"
)

// New returns an instantiated novaexchange struct
func New(apiKey, apiSecret string) *NovaExchange {
	client := NewClient(apiKey, apiSecret)
	return &NovaExchange{client}
}

// NewWithCustomHttpClient returns an instantiated novaexchange struct with custom http client
func NewWithCustomHttpClient(apiKey, apiSecret string, httpClient *http.Client) *NovaExchange {
	client := NewClientWithCustomHttpConfig(apiKey, apiSecret, httpClient)
	return &NovaExchange{client}
}

// NewWithCustomTimeout returns an instantiated novaexchange struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *NovaExchange {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &NovaExchange{client}
}

// handleErr gets JSON response from Novaexchange API en deal with error
func handleErr(r jsonResponse) error {
	if r.Status != "success" {
		return errors.New(r.Message)
	}
	return nil
}

// novaexchange represent a novaexchange client
type NovaExchange struct {
	client *client
}

// set enable/disable http request/response dump
func (c *NovaExchange) SetDebug(enable bool) {
	c.client.debug = enable
}


// Account

// GetBalances is used to retrieve all balances from your account
func (b *NovaExchange) GetBalances() (balances []Balance, err error) {
	r, err := b.client.do("GET", "private/getbalances/", "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result["balances"], &balances)
	return
}

// Getbalance is used to retrieve the balance from your account for a specific currency.
// currency: a string literal for the currency (ex: LTC)
func (b *NovaExchange) GetBalance(currency string) (balance Balance, err error) {
	r, err := b.client.do("GET", "private/getbalance/"+strings.ToUpper(currency)+"/", "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result["balances"], &balance)
	return
}

// GetDepositAddress is sed to generate or retrieve an address for a specific currency.
// currency a string literal for the currency (ie. BTC)
func (b *NovaExchange) GetDepositAddress(currency string) (address string, err error) {
	r, err := b.client.do("GET", "private/getdepositaddress/"+strings.ToUpper(currency)+"/", "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result["address"], &address)
	return
}

// GetOrderHistory used to retrieve your order history.
// market string literal for the market (ie. BTC-LTC). If set to "all", will return for all market
func (b *NovaExchange) GetOrderHistory(page int) (orders []Order, err error) {
	ressource := "private/tradehistory/"
	if page != 0 {
		ressource += fmt.Sprintf("?page=%d", page)
	}
	r, err := b.client.do("POST", ressource, "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result["items"], &orders)
	return
}

// GetWithdrawalHistory is used to retrieve your withdrawal history
// currency string a string literal for the currency (ie. BTC). If set to "all", will return for all currencies
func (b *NovaExchange) GetWithdrawalHistory(page int) (withdrawals []Withdrawal, err error) {
	ressource := "private/getwithdrawalhistory/"
	if page != 0 {
		ressource += fmt.Sprintf("?page=%d", page)
	}
	r, err := b.client.do("POST", ressource, "", true)
	if err != nil {
		return
	}
	log.Print(string(r))
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result["items"], &withdrawals)
	return
}

// GetDepositHistory is used to retrieve your deposit history
// currency string a string literal for the currency (ie. BTC). If set to "all", will return for all currencies
func (b *NovaExchange) GetDepositHistory(page int) (deposits []Deposit, err error) {
	ressource := "private/getdeposithistory/"
	if page != 0 {
		ressource += fmt.Sprintf("?page=%d", page)
	}
	r, err := b.client.do("POST", ressource, "", true)
	if err != nil {
		return
	}
	log.Print(string(r))
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result["items"], &deposits)
	return
}
