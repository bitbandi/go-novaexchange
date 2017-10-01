package novaexchange

type Deposit struct {
	Status     string  `json:"status"`
	CurrencyId int64   `json:"currencyid"`
	Currency   string  `json:"currency"`
	Amount     float64 `json:"tx_amount,string"`
	Date       jTime   `json:"time_seen"`
	TxId       string  `json:"tx_txid"`
	Address    string  `json:"tx_address"`
}
