package parameters

import (
	"fmt"
)

type NewOrder struct {
	Symbol           string  `json:"symbol"`
	Side             string  `json:"side"`
	Type             string  `json:"type"`
	TimeInForce      string  `json:"timeInForce"`      // optional
	Quantity         float64 `json:"quantity"`         // optional
	QuoteOrderQty    float64 `json:"quoteOrderQty"`    // optional
	Price            float64 `json:"price"`            // optional
	NewClientOrderId string  `json:"newClientOrderId"` // optional
	StrategyId       int64   `json:"strategyId"`       // optional
	StrategyType     int64   `json:"strategyType"`     // optional
	StopPrice        float64 `json:"stopPrice"`        // optional
	TrailingDelta    float64 `json:"trailingDelta"`    // optional
	IcebergQty       float64 `json:"icebergQty"`       // optional
	NewOrderRespType string  `json:"newOrderRespType"` // optional
	RecvWindow       int64   `json:"recvWindow"`       // optional
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
