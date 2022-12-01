package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#query-margin-priceindex-market_data
	EndpointMarginAccountPriceIndex = "/sapi/v1/margin/priceIndex"
)

type MarginAccountPriceIndexParam struct {
	Symbol string
}

type MarginAccountPriceIndexResp struct {
	CalcTime int64  `json:"calcTime"`
	Price    string `json:"price"`
	Symbol   string `json:"symbol"`
}

func (ex *BinanceRest) MarginAccountPriceIndex(parm MarginAccountPriceIndexParam) MarginAccountPriceIndexResp {
	r := &Request{
		method:   http.MethodGet,
		endpoint: EndpointMarginAccountPriceIndex,
		secType:  secTypeAPIKey,
	}
	m := setMarginAccountPriceIndexParams(parm)
	r.setParams(m)
	data, err := ex.callAPI(r)

	if err != nil {
		log.Printf(`
 				{
 					"Status" : "Error",
 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
 					"File": "MarginAccountPriceIndex.go",
 					"Functions" : "(ex *BinanceRest) MarginAccountPriceIndex(parm MarginAccountPriceIndexParam) MarginAccountPriceIndexResp",
 					"Function where err" : "ex.callAPI",
 					"Exchange" : "Binance",
 					"Error" : %s
 				}`, err)
		log.Fatal()
	}

	var marginAccountPriceIndex MarginAccountPriceIndexResp
	_ = json.Unmarshal(data, &marginAccountPriceIndex)
	return marginAccountPriceIndex
}

func setMarginAccountPriceIndexParams(parm MarginAccountPriceIndexParam) params {
	m := params{
		"symbol": parm.Symbol,
	}
	return m
}
