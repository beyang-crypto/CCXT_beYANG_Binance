package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#account-information-user_data
	EndpointAccountInformation = "/api/v3/account"
)

type AccountInformationParam struct {
	RecvWindow *int64 // optional
}

type AccountInformationResp struct {
	MakerCommission  int
	TakerCommission  int
	BuyerCommission  int
	SellerCommission int
	CanTrade         bool
	CanWithdraw      bool
	CanDeposit       bool
	Brokered         bool
	UpdateTime       int
	AccountType      string
	Balances         []struct {
		Asset  string
		Free   string
		Locked string
	}
	Permissions []string
}

func (ex *BinanceRest) AccountInformation(parm AccountInformationParam) AccountInformationResp {
	r := &Request{
		method:   http.MethodGet,
		endpoint: EndpointAccountInformation,
		secType:  secTypeSigned,
	}
	m := setAccountInformationParams(parm)
	r.setParams(m)

	data, err := ex.callAPI(r)

	if err != nil {
		log.Printf(`
 				{
 					"Status" : "Error",
 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
 					"File": "AccountInformation.go",
 					"Functions" : "(ex *BinanceRest) AccountInformation(parm AccountInformationParam) AccountInformationResp",
 					"Function where err" : "ex.callAPI",
 					"Exchange" : "Binance",
 					"Error" : %s
 				}`, err)
		log.Fatal()
	}
	var accountInformation AccountInformationResp
	_ = json.Unmarshal(data, &accountInformation)
	return accountInformation
}

func setAccountInformationParams(parm AccountInformationParam) params {
	m := params{
		"recvWindow": parm.RecvWindow,
	}
	return m
}
