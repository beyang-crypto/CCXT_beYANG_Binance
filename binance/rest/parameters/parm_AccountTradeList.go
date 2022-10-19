package parameters

import "fmt"

type AccountTradeList struct {
	Symbol     string `json:"symbol"`
	OrderId    int64  `json:"orderId"`
	StartTime  int64  `json:"startTime"`
	EndTime    int64  `json:"endTime"`
	FromId     int64  `json:"fromId"`
	Limit      int64  `json:"limit"`
	RecvWindow int64  `json:"recvWindow"`
}

func BinanceToAccountTradeListParms(data interface{}) (AccountTradeList, bool) {
	atl, ok := data.(AccountTradeList)
	return atl, ok
}

func BinanceParmAccountTradeListToString(parm AccountTradeList) string {
	par := ""
	par += checkSymbol(parm.Symbol)
	par += checkOrderId(parm.OrderId)
	par += checkStartTime(parm.StartTime)
	par += checkEndTime(parm.EndTime)
	par += checkFromId(parm.FromId)
	par += checkLimit(parm.Limit)
	par += checkRecvWindow(parm.RecvWindow)
	par += fmt.Sprintf("timestamp=%d", getTimestamp())
	return par
}
