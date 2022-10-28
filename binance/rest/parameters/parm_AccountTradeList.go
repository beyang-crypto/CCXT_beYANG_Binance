package parameters

import "fmt"

type AccountTradeList struct {
	Symbol     string
	OrderId    int64
	StartTime  int64
	EndTime    int64
	FromId     int64
	Limit      int64
	RecvWindow int64
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
