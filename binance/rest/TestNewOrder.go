package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

// https://binance-docs.github.io/apidocs/spot/en/#test-new-order-trade
const EndpointTestNewOrder = "/api/v3/order/test"

type TestNewOrderParam struct {
	Symbol           string
	Side             string
	Type             string
	TimeInForce      *string  // optional
	Quantity         *float64 // optional
	QuoteOrderQty    *float64 // optional
	Price            *float64 // optional
	NewClientOrderId *string  // optional
	StrategyId       *int64   // optional
	StrategyType     *int64   // optional
	StopPrice        *float64 // optional
	TrailingDelta    *float64 // optional
	IcebergQty       *float64 // optional
	NewOrderRespType *string  // optional
	RecvWindow       *int64   // optional
}

type TestNewOrderResp struct {
}

func (ex *BinanceRest) TestNewOrder(parm TestNewOrderParam) TestNewOrderResp {
	r := &Request{
		method:   http.MethodPost,
		endpoint: EndpointTestNewOrder,
		secType:  SecTypeSigned,
	}
	m := setTestNewOrderParams(parm)
	r.setParams(m)

	data, err := ex.callAPI(r)

	if err != nil {
		log.Printf("%v", err)
	}
	var testNewOrder TestNewOrderResp
	_ = json.Unmarshal(data, &testNewOrder)
	return testNewOrder
}

func setTestNewOrderParams(parm TestNewOrderParam) Params {
	m := Params{
		"symbol":     parm.Symbol,
		"side":       parm.Side,
		"type":       parm.Type,
		"recvWindow": parm.RecvWindow,
	}
	if parm.Quantity != nil {
		m["quantity"] = *parm.Quantity
	}
	if parm.QuoteOrderQty != nil {
		m["quoteOrderQty"] = *parm.QuoteOrderQty
	}
	if parm.TimeInForce != nil {
		m["timeInForce"] = *parm.TimeInForce
	}
	if parm.Price != nil {
		m["price"] = *parm.Price
	}
	if parm.NewClientOrderId != nil {
		m["newClientOrderId"] = *parm.NewClientOrderId
	}
	if parm.StopPrice != nil {
		m["stopPrice"] = *parm.StopPrice
	}
	if parm.TrailingDelta != nil {
		m["trailingDelta"] = *parm.TrailingDelta
	}
	if parm.IcebergQty != nil {
		m["icebergQty"] = *parm.IcebergQty
	}
	if parm.NewOrderRespType != nil {
		m["newOrderRespType"] = *parm.NewOrderRespType
	}
	return m
}
