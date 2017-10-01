package novaexchange

type Balance struct {
	Currency     string  `json:"currency"`
	CurrencyId   int     `json:"currencyid"`
	CurrencyName string  `json:"currencyname"`
	Balance      float64 `json:"amount,string"`
	Trades       float64 `json:"amount_trades,string"`
	LockBox      float64 `json:"amount_lockbox,string"`
	Total        float64 `json:"amount_total,string"`
}
