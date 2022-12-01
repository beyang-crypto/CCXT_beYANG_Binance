package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#query-margin-account-39-s-order-user_data
	EndpointMarginAccountQueryOrder = "/sapi/v1/margin/order"
)

type MarginAccountQueryOrderParam struct {
	Symbol            string
	IsIsolated        *string // optional
	OrderId           *int64  // optional
	OrigClientOrderId *string // optional
	RecvWindow        *int64  // optional
}

type MarginAccountQueryOrderResp struct {
	ClientOrderID       string `json:"clientOrderId"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	ExecutedQty         string `json:"executedQty"`
	IcebergQty          string `json:"icebergQty"`
	IsWorking           bool   `json:"isWorking"`
	OrderID             int    `json:"orderId"`
	OrigQty             string `json:"origQty"`
	Price               string `json:"price"`
	Side                string `json:"side"`
	Status              string `json:"status"`
	StopPrice           string `json:"stopPrice"`
	Symbol              string `json:"symbol"`
	IsIsolated          bool   `json:"isIsolated"`
	Time                int64  `json:"time"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	UpdateTime          int64  `json:"updateTime"`
}

func (ex *BinanceRest) MarginAccountQueryOrder(parm MarginAccountQueryOrderParam) MarginAccountQueryOrderResp {
	r := &Request{
		method:   http.MethodGet,
		endpoint: EndpointMarginAccountQueryOrder,
		secType:  secTypeSigned,
	}

	m := setMarginAccountQueryOrderParams(parm)
	r.setParams(m)

	data, err := ex.callAPI(r)
	if err != nil {
		log.Printf(`
 				{
 					"Status" : "Error",
 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
 					"File": "MarginAccountQueryOrder.go",
 					"Functions" : "(ex *BinanceRest) MarginAccountQueryOrder(parm MarginAccountQueryOrderParam) MarginAccountQueryOrderResp",
 					"Function where err" : "ex.callAPI",
 					"Exchange" : "Binance",
 					"Error" : %s
 				}`, err)
		log.Fatal()
	}

	var marginAccountQueryOrder MarginAccountQueryOrderResp
	_ = json.Unmarshal(data, &marginAccountQueryOrder)
	return marginAccountQueryOrder
}

func setMarginAccountQueryOrderParams(parm MarginAccountQueryOrderParam) params {
	m := params{
		"symbol": parm.Symbol,
	}
	if parm.IsIsolated != nil {
		m["isIsolated"] = *parm.IsIsolated
	}
	if parm.OrderId != nil {
		m["orderId"] = *&parm.OrderId
	}
	if parm.OrigClientOrderId != nil {
		m["origClientOrderId"] = *&parm.OrigClientOrderId
	}
	if parm.RecvWindow != nil {
		m["recvWindow"] = *&parm.RecvWindow
	}
	return m
}
