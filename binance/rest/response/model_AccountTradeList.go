package response

type AccountTradeList []struct {
	Symbol          string
	ID              int
	OrderID         int
	OrderListID     int
	Price           string
	Qty             string
	QuoteQty        string
	Commission      string
	CommissionAsset string
	Time            int64
	IsBuyer         bool
	IsMaker         bool
	IsBestMatch     bool
}

func BinanceToAccountTradeList(data interface{}) (AccountTradeList, bool) {
	atl, ok := data.(AccountTradeList)
	return atl, ok
}
