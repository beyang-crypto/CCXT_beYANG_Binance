package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#margin-account-new-order-trade
	EndpointMarginAccountNewOrder = "/sapi/v1/margin/order"
)

type MarginAccountNewOrderParam struct {
	Symbol           string
	IsIsolated       *string // optional
	Side             OrderSideType
	Type             OrderType
	Quantity         *float64          // optional
	QuoteOrderQty    *float64          // optional
	Price            *float64          // optional
	StopPrice        *float64          // optional
	NewClientOrderId *string           // optional
	IcebergQty       *float64          // optional
	NewOrderRespType *NewOrderRespType // optional
	SideEffectType   *SideEffectType   // optional
	TimeInForce      *TimeInForceType  // optional
	RecvWindow       *int64            // optional
}
type MarginAccountNewOrderResp struct {
	Symbol                string `json:"symbol"`
	OrderID               int    `json:"orderId"`
	ClientOrderID         string `json:"clientOrderId"`
	TransactTime          int64  `json:"transactTime"`
	Price                 string `json:"price"`
	OrigQty               string `json:"origQty"`
	ExecutedQty           string `json:"executedQty"`
	CummulativeQuoteQty   string `json:"cummulativeQuoteQty"`
	Status                string `json:"status"`
	TimeInForce           string `json:"timeInForce"`
	Type                  string `json:"type"`
	Side                  string `json:"side"`
	MarginBuyBorrowAmount int    `json:"marginBuyBorrowAmount"`
	MarginBuyBorrowAsset  string `json:"marginBuyBorrowAsset"`
	IsIsolated            bool   `json:"isIsolated"`
	Fills                 []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
	} `json:"fills"`
}

func (ex *BinanceRest) MarginAccountNewOrder(parm MarginAccountNewOrderParam) MarginAccountNewOrderResp {
	r := &Request{
		method:   http.MethodPost,
		endpoint: EndpointMarginAccountNewOrder,
		secType:  secTypeSigned,
	}

	m := setMarginAccountNewOrderParams(parm)
	r.setParams(m)

	data, err := ex.callAPI(r)

	if err != nil {
		log.Printf(`
 				{
 					"Status" : "Error",
 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
 					"File": "MarginAccountNewOrder.go",
 					"Functions" : "(ex *BinanceRest) MarginAccountNewOrder(parm MarginAccountNewOrderParam) MarginAccountNewOrderResp",
 					"Function where err" : "ex.callAPI",
 					"Exchange" : "Binance",
 					"Error" : %s
 				}`, err)
		log.Fatal()
	}

	var marginAccountNewOrder MarginAccountNewOrderResp
	_ = json.Unmarshal(data, &marginAccountNewOrder)
	return marginAccountNewOrder
}

func setMarginAccountNewOrderParams(parm MarginAccountNewOrderParam) params {
	m := params{
		"symbol": parm.Symbol,
		"side":   parm.Side,
		"type":   parm.Type,
	}
	if parm.IsIsolated != nil {
		m["isIsolated"] = *parm.IsIsolated
	}
	if parm.Quantity != nil {
		m["quantity"] = *parm.Quantity
	}
	if parm.QuoteOrderQty != nil {
		m["quoteOrderQty"] = *parm.QuoteOrderQty
	}
	if parm.Price != nil {
		m["price"] = *parm.Price
	}
	if parm.StopPrice != nil {
		m["stopPrice"] = *parm.StopPrice
	}
	if parm.NewClientOrderId != nil {
		m["newClientOrderId"] = *parm.NewClientOrderId
	}
	if parm.IcebergQty != nil {
		m["icebergQty"] = *parm.IcebergQty
	}
	if parm.NewOrderRespType != nil {
		m["newOrderRespType"] = *parm.NewOrderRespType
	}
	if parm.TimeInForce != nil {
		m["sideEffectType"] = *parm.SideEffectType
	}
	if parm.TimeInForce != nil {
		m["timeInForce"] = *parm.TimeInForce
	}
	if parm.RecvWindow != nil {
		m["recvWindow"] = *&parm.RecvWindow
	}
	return m
}
