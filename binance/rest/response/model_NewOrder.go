package response

type NewOrder struct {
	Symbol              string
	OrderID             int
	OrderListID         int
	ClientOrderID       string
	TransactTime        int64
	Price               string
	OrigQty             string
	ExecutedQty         string
	CummulativeQuoteQty string
	Status              string
	TimeInForce         string
	Type                string
	Side                string
	StrategyType        int
	Fills               []Fills
}
type Fills struct {
	Price           string
	Qty             string
	Commission      string
	CommissionAsset string
	TradeID         int
}

func BinanceToNewOrder(data interface{}) (TestNewOrder, bool) {
	tno, ok := data.(TestNewOrder)
	return tno, ok
}
