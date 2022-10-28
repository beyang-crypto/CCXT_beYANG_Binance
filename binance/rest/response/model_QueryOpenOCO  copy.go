package response

type QueryOpenOCO []struct {
	OrderListID       int
	ContingencyType   string
	ListStatusType    string
	ListOrderStatus   string
	ListClientOrderID string
	TransactionTime   int64
	Symbol            string
	Orders            []Orders
}
type Orders struct {
	Symbol        string
	OrderID       int
	ClientOrderID string
}

func BinanceToQueryOpenOCO(data interface{}) (QueryOpenOCO, bool) {
	qooco, ok := data.(QueryOpenOCO)
	return qooco, ok
}
