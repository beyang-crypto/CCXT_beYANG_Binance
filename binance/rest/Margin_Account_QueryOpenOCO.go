package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#query-open-oco-user_data
	EndpointMarginQueryOpenOCO = "/sapi/v1/margin/openOrderList"
)

type MarginQueryOpenOCOParam struct {
	IsIsolated *string // optional
	Symbol     *string // optional
	RecvWindow *int64  // optional
}

type MarginQueryOpenOCOResp []struct {
	OrderListID       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderID string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderID       int    `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
	} `json:"orders"`
}

func (ex *BinanceRest) MarginQueryOpenOCO(parm MarginQueryOpenOCOParam) MarginQueryOpenOCOResp {
	r := &Request{
		method:   http.MethodGet,
		endpoint: EndpointMarginQueryOpenOCO,
		secType:  secTypeSigned,
	}

	m := setMarginQueryOpenOCOParams(parm)
	r.setParams(m)

	data, err := ex.callAPI(r)
	if err != nil {
		log.Printf("%v", err)
	}

	var marginQueryOpenOCOResp MarginQueryOpenOCOResp
	_ = json.Unmarshal(data, &marginQueryOpenOCOResp)
	return marginQueryOpenOCOResp
}

func setMarginQueryOpenOCOParams(parm MarginQueryOpenOCOParam) params {
	m := params{}
	if parm.IsIsolated != nil {
		m["isIsolated"] = *parm.IsIsolated
	}
	if parm.Symbol != nil {
		m["symbol"] = *parm.Symbol
	}
	if parm.RecvWindow != nil {
		m["recvWindow"] = *&parm.RecvWindow
	}
	return m
}
