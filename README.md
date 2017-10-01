go-novaexchange
==========

go-novaexchange is an implementation of the Novaexchange API (public and private) in Golang.

Based off of https://github.com/toorop/go-bittrex/

## Import
	import "github.com/bitbandi/go-novaexchange"
	
## Usage
~~~ go
package main

import (
	"fmt"
	"github.com/bitbandi/go-novaexchange"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// Novaexchange client
	novaexchange := novaexchange.New(API_KEY, API_SECRET)

	// Get balances
	balances, err := novaexchange.GetBalances()
	fmt.Println(err, balances)
}
~~~

See ["Examples" folder for more... examples](https://github.com/bitbandi/go-novaexchange/blob/master/examples/novaexchange.go)
