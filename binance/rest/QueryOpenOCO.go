package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#query-open-oco-user_data
	EndpointQueryOpenOCO = "/api/v3/openOrderList"
)

type QueryOpenOCOParam struct {
	RecvWindow *int64 // optional
}

type QueryOpenOCOResp []struct {
	OrderListID       int
	ContingencyType   string
	ListStatusType    string
	ListOrderStatus   string
	ListClientOrderID string
	TransactionTime   int64
	Symbol            string
	Orders            []struct {
		Symbol        string
		OrderID       int
		ClientOrderID string
	}
}

func (ex *BinanceRest) QueryOpenOCO(parm QueryOpenOCOParam) QueryOpenOCOResp {
	r := &Request{
		method:   http.MethodGet,
		endpoint: EndpointUserAsset,
		secType:  SecTypeSigned,
	}

	m := setQueryOpenOCOParams(parm)
	r.setParams(m)

	data, err := ex.callAPI(r)
	if err != nil {
		log.Printf("%v", err)
	}

	var queryOpenOCO QueryOpenOCOResp
	_ = json.Unmarshal(data, &queryOpenOCO)
	return queryOpenOCO
}

func setQueryOpenOCOParams(parm QueryOpenOCOParam) Params {
	m := Params{
		"recvWindow": parm.RecvWindow,
	}
	return m
}
