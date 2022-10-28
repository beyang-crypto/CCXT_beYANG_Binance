package response

type QueryOrder struct {
	Symbol              string
	OrderID             int
	OrderListID         int
	ClientOrderID       string
	Price               string
	OrigQty             string
	ExecutedQty         string
	CummulativeQuoteQty string
	Status              string
	TimeInForce         string
	Type                string
	Side                string
	StopPrice           string
	IcebergQty          string
	Time                int64
	UpdateTime          int64
	IsWorking           bool
	OrigQuoteOrderQty   string
}

func BinanceToQueryOrder(data interface{}) (QueryOrder, bool) {
	qo, ok := data.(QueryOrder)
	return qo, ok
}
