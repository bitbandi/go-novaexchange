package novaexchange

type Withdrawal struct {
	Status     string  `json:"status"`
	Currency   string  `json:"currency"`
	CurrencyId int64   `json:"currencyid"`
	Requested  jTime   `json:"time_requested"`
	Send       jTime   `json:"time_sent"`
	Address    string  `json:"tx_address"`
	Amount     float64 `json:"tx_amount,string"`
	TxId       string  `json:"tx_txid"`
}
