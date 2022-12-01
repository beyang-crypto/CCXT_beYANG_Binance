package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#query-order-user_data
	EndpointQueryOrder = "/api/v3/order"
)

type QueryOrderParam struct {
	Symbol            string
	OrderId           *int64  // optional
	OrigClientOrderId *string // optional
	RecvWindow        *int64  // optional
}

type QueryOrderResp struct {
	Symbol              string
	OrderID             int
	OrderListID         int
	ClientOrderID       string
	Price               string
	OrigQty             string
	ExecutedQty         string
	CummulativeQuoteQty string
	Status              string
	TimeInForce         string
	Type                string
	Side                string
	StopPrice           string
	IcebergQty          string
	Time                int64
	UpdateTime          int64
	IsWorking           bool
	OrigQuoteOrderQty   string
}

func (ex *BinanceRest) QueryOrder(parm QueryOrderParam) QueryOrderResp {
	r := &Request{
		method:   http.MethodGet,
		endpoint: EndpointQueryOrder,
		secType:  secTypeSigned,
	}

	m := setQueryOrderParams(parm)
	r.setParams(m)

	data, err := ex.callAPI(r)
	if err != nil {
		log.Printf(`
 				{
 					"Status" : "Error",
 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
 					"File": "QueryOrder.go",
 					"Functions" : "(ex *BinanceRest) QueryOrder(parm QueryOrderParam) QueryOrderResp",
 					"Function where err" : "ex.callAPI",
 					"Exchange" : "Binance",
 					"Error" : %s
 				}`, err)
		log.Fatal()
	}

	var queryOrder QueryOrderResp
	_ = json.Unmarshal(data, &queryOrder)
	return queryOrder
}

func setQueryOrderParams(parm QueryOrderParam) params {
	m := params{
		"symbol": parm.Symbol,
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
