package parameters

import (
	"fmt"
)

type NewOrder struct {
	Symbol           string
	Side             string
	Type             string
	TimeInForce      string  // optional
	Quantity         float64 // optional
	QuoteOrderQty    float64 // optional
	Price            float64 // optional
	NewClientOrderId string  // optional
	StrategyId       int64   // optional
	StrategyType     int64   // optional
	StopPrice        float64 // optional
	TrailingDelta    float64 // optional
	IcebergQty       float64 // optional
	NewOrderRespType string  // optional
	RecvWindow       int64   // optional
}

func BinanceParmsToNewOrder(data interface{}) (TestNewOrder, bool) {
	tno, ok := data.(TestNewOrder)
	return tno, ok
}

func BinanceParmNewOrderToString(parm TestNewOrder) string {
	par := ""
	par += checkSymbol(parm.Symbol)
	par += checkSide(parm.Side)
	par += checkType(parm.Type)
	par += checkTimeInForce(parm.TimeInForce)
	par += checkQuantity(parm.Quantity)
	par += checkQuoteOrderQty(parm.QuoteOrderQty)
	par += checkPrice(parm.Price)
	par += checkRecvWindow(parm.RecvWindow)
	par += fmt.Sprintf("timestamp=%d", getTimestamp())
	return par
}
