package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#account-trade-list-user_data
	EndpointAccountTradeList = "/api/v3/myTrades"
)

type AccountTradeListParam struct {
	Symbol     string
	OrderId    *int64 // optional
	StartTime  *int64 // optional
	EndTime    *int64 // optional
	FromId     *int64 // optional
	Limit      *int64 // optional
	RecvWindow *int64 // optional
}

type AccountTradeListResp []struct {
	Symbol          string
	ID              int
	OrderID         int
	OrderListID     int
	Price           string
	Qty             string
	QuoteQty        string
	Commission      string
	CommissionAsset string
	Time            int64
	IsBuyer         bool
	IsMaker         bool
	IsBestMatch     bool
}

func (ex *BinanceRest) AccountTradeList(parm AccountTradeListParam) AccountTradeListResp {
	r := &Request{
		method:   http.MethodGet,
		endpoint: EndpointAccountTradeList,
		secType:  secTypeSigned,
	}

	m := setAccountTradeListParams(parm)
	r.setParams(m)

	data, err := ex.callAPI(r)

	if err != nil {
		log.Printf(`
 				{
 					"Status" : "Error",
 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
 					"File": "AccountTradeList.go",
 					"Functions" : "(ex *BinanceRest) AccountTradeList(parm AccountTradeListParam) AccountTradeListResp",
 					"Function where err" : "ex.callAPI",
 					"Exchange" : "Binance",
 					"Error" : %s
 				}`, err)
		log.Fatal()
	}

	var accountTradeList AccountTradeListResp
	_ = json.Unmarshal(data, &accountTradeList)
	return accountTradeList
}

func setAccountTradeListParams(parm AccountTradeListParam) params {
	m := params{
		"symbol": parm.Symbol,
	}
	if parm.OrderId != nil {
		m["orderId"] = *&parm.OrderId
	}
	if parm.StartTime != nil {
		m["startTime"] = *&parm.OrderId
	}
	if parm.EndTime != nil {
		m["endTime"] = *&parm.OrderId
	}
	if parm.FromId != nil {
		m["fromId"] = *&parm.OrderId
	}
	if parm.Limit != nil {
		m["limit"] = *&parm.OrderId
	}
	if parm.RecvWindow != nil {
		m["recvWindow"] = *&parm.OrderId
	}
	return m
}
