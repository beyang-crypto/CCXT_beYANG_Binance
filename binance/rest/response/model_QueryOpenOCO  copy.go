package response

type QueryOpenOCO []struct {
	OrderListID       int      `json:"orderListId"`
	ContingencyType   string   `json:"contingencyType"`
	ListStatusType    string   `json:"listStatusType"`
	ListOrderStatus   string   `json:"listOrderStatus"`
	ListClientOrderID string   `json:"listClientOrderId"`
	TransactionTime   int64    `json:"transactionTime"`
	Symbol            string   `json:"symbol"`
	Orders            []Orders `json:"orders"`
}
type Orders struct {
	Symbol        string `json:"symbol"`
	OrderID       int    `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}

func BinanceToQueryOpenOCO(data interface{}) (QueryOpenOCO, bool) {
	qooco, ok := data.(QueryOpenOCO)
	return qooco, ok
}
