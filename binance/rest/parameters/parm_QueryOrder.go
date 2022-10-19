package parameters

import (
	"fmt"
)

type QueryOrder struct {
	Symbol            string `json:"symbol"`
	OrderId           int64  `json:"orderId"`           // optional
	OrigClientOrderId string `json:"origClientOrderId"` // optional
	RecvWindow        int64  `json:"recvWindow"`        // optional
}

func BinanceParmsToQueryOrder(data interface{}) (TestNewOrder, bool) {
	tno, ok := data.(TestNewOrder)
	return tno, ok
}

func BinanceParmTestQueryOrderToString(parm QueryOrder) string {
	par := ""
	par += checkSymbol(parm.Symbol)
	par += checkOrderId(parm.OrderId)
	par += checkOrigClientOrderId(parm.OrigClientOrderId)
	par += checkRecvWindow(parm.RecvWindow)
	par += fmt.Sprintf("timestamp=%d", getTimestamp())
	return par
}
