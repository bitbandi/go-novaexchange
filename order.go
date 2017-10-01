package novaexchange

type Order struct {
	Id           int     `json:"tradeid"`
	TradeType    string  `json:"tradetype"`
	Price        float64 `json:"price,string"`
	ToAmount     float64 `json:"toamount,string"`
	FromAmount   float64 `json:"fromamount,string"`
	Fee          float64 `json:"fee,string"`
	TradeTime    jTime   `json:"trade_time"`
	ToCurrency   string  `json:"tocurrency"`
	FromCurrency string  `json:"fromcurrency"`
	BaseCurrency string  `json:"basecurrency"`
	OrderId      int     `json:"orig_orderid"`
}
