package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#new-order-trade
	EndpointNewOrder = "/api/v3/order"
)

type NewOrderParam struct {
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

type NewOrderResp struct {
	Symbol              string
	OrderID             int
	OrderListID         int
	ClientOrderID       string
	TransactTime        int64
	Price               string
	OrigQty             string
	ExecutedQty         string
	CummulativeQuoteQty string
	Status              string
	TimeInForce         string
	Type                string
	Side                string
	StrategyType        int
	Fills               []struct {
		Price           string
		Qty             string
		Commission      string
		CommissionAsset string
		TradeID         int
	}
}

func (ex *BinanceRest) NewOrder(parm NewOrderParam) NewOrderResp {
	r := &Request{
		method:   http.MethodPost,
		endpoint: EndpointNewOrder,
		secType:  SecTypeSigned,
	}

	m := setNewOrder(parm)

	r.setParams(m)

	data, err := ex.callAPI(r)

	if err != nil {
		log.Printf("%v", err)
	}

	var newOrder NewOrderResp
	_ = json.Unmarshal(data, &newOrder)
	return newOrder
}

func setNewOrder(parm NewOrderParam) Params {
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
