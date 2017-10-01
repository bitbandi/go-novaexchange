package main

import (
	"fmt"

	"github.com/bitbandi/go-novaexchange"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	// novaexchange client
	novaexchange := novaexchange.New(API_KEY, API_SECRET)

	// GetBalances
	balances, _ := novaexchange.GetBalances()
	fmt.Println(len(balances))

	for i, _ := range balances {
		if balances[i].Currency == "BTC" {
			fmt.Println(balances[i].Balance)
		}
	}

}
