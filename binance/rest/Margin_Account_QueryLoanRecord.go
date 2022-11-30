package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#query-loan-record-user_data
	EndpointQueryLoanRecord = "/sapi/v1/margin/loan"
)

type QueryLoanRecordParam struct {
	Asset          string
	IsolatedSymbol *string  // optional
	TxId           *float64 // optional
	StartTime      *float64 // optional
	EndTime        *float64 // optional
	Current        *float64 // optional
	Size           *float64 // optional
	Archived       *string  // optional
	RecvWindow     *int64   // optional
}

type QueryLoanRecordResp struct {
	Rows []struct {
		IsolatedSymbol string `json:"isolatedSymbol"`
		TxID           int64  `json:"txId"`
		Asset          string `json:"asset"`
		Principal      string `json:"principal"`
		Timestamp      int64  `json:"timestamp"`
		Status         string `json:"status"`
	} `json:"rows"`
	Total int `json:"total"`
}

func (ex *BinanceRest) QueryLoanRecord(parm QueryLoanRecordParam) QueryLoanRecordResp {
	r := &Request{
		method:   http.MethodGet,
		endpoint: EndpointQueryLoanRecord,
		secType:  secTypeSigned,
	}

	m := setQueryLoanRecordParams(parm)

	r.setParams(m)

	data, err := ex.callAPI(r)

	if err != nil {
		log.Printf("%v", err)
	}

	var queryLoanRecord QueryLoanRecordResp
	_ = json.Unmarshal(data, &queryLoanRecord)
	return queryLoanRecord
}

func setQueryLoanRecordParams(parm QueryLoanRecordParam) params {
	m := params{
		"asset": parm.Asset,
	}
	if parm.IsolatedSymbol != nil {
		m["isolatedSymbol"] = *parm.IsolatedSymbol
	}
	if parm.TxId != nil {
		m["txId"] = *parm.TxId
	}
	if parm.StartTime != nil {
		m["startTime"] = *parm.StartTime
	}
	if parm.EndTime != nil {
		m["endTime"] = *parm.EndTime
	}
	if parm.TxId != nil {
		m["current"] = *parm.Current
	}
	if parm.TxId != nil {
		m["size"] = *parm.Size
	}
	if parm.TxId != nil {
		m["archived"] = *parm.Archived
	}
	if parm.RecvWindow != nil {
		m["recvWindow"] = *&parm.RecvWindow
	}
	return m
}
