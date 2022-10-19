package response

type NewOrder struct {
	Symbol              string  `json:"symbol"`
	OrderID             int     `json:"orderId"`
	OrderListID         int     `json:"orderListId"`
	ClientOrderID       string  `json:"clientOrderId"`
	TransactTime        int64   `json:"transactTime"`
	Price               string  `json:"price"`
	OrigQty             string  `json:"origQty"`
	ExecutedQty         string  `json:"executedQty"`
	CummulativeQuoteQty string  `json:"cummulativeQuoteQty"`
	Status              string  `json:"status"`
	TimeInForce         string  `json:"timeInForce"`
	Type                string  `json:"type"`
	Side                string  `json:"side"`
	StrategyType        int     `json:"strategyType"`
	Fills               []Fills `json:"fills"`
}
type Fills struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	TradeID         int    `json:"tradeId"`
}

func BinanceToNewOrder(data interface{}) (TestNewOrder, bool) {
	tno, ok := data.(TestNewOrder)
	return tno, ok
}
