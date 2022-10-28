package rest

import (
	"log"

	"github.com/TestingAccMar/CCXT_beYANG_Binance/binance/rest/parameters"
	"github.com/TestingAccMar/CCXT_beYANG_Binance/binance/rest/response"
	"github.com/goccy/go-json"
)

func (ex *BinanceRest) Get(endpoint string, parms interface{}) interface{} {

	parmsForEndpoint := false
	withAuth := false
	par := ""
	switch endpoint {
	case EndpointAccountInformation:
		parm, ok := parameters.BinanceParmsToAccountInformation(parms)
		if ok {
			par = parameters.BinanceParmAccountInformationToString(parm)
			withAuth = true
		} else {
			parmsForEndpoint = true
		}
	case EndpointAccountTradeList:
		parm, ok := parameters.BinanceToAccountTradeListParms(parms)
		if ok {
			par = parameters.BinanceParmAccountTradeListToString(parm)
			withAuth = true
		} else {
			parmsForEndpoint = true
		}
	case EndpointQueryOpenOCO:
		parm, ok := parameters.BinanceParmsToQueryOpenOCO(parms)
		if ok {
			par = parameters.BinanceParmQueryOpenOCOToString(parm)
			withAuth = true
		} else {
			parmsForEndpoint = true
		}
	case EndpointExchangeInfo:
		parm, ok := parameters.BinanceToExchangeInformationParms(parms)
		if ok {
			par = parameters.BinanceParmExchangeInformationToString(parm)
			withAuth = false
		} else {
			parmsForEndpoint = true
		}
	default:
		log.Printf(`
			{
				"Status" : "Error",
				"Path to file" : "CCXT_beYANG_Binance/binance/rest",
				"File": "get.go",
				"Functions" : "(ex *BinanceRest) Get(endpoint string, parms interface{})",
				"Exchange" : "Binance",
				"Error" : Unable to connect to this endpoint,
			}`)
		log.Fatal()
	}
	// Если параметр не соответствует конечной точке
	if parmsForEndpoint {
		log.Printf(`
			{
				"Status" : "Error",
				"Path to file" : "CCXT_beYANG_Binance/binance/rest",
				"File": "get.go",
				"Functions" : "(ex *BinanceRest) Get(endpoint string, parms interface{})",
				"Exchange" : "Binance",
				"Error" : the parameters do not match the endpoint(%s),
			}`, endpoint)
		log.Fatal()
	}

	var data []byte

	var returnInterface interface{}

	if withAuth {
		par += "&signature=" + ex.GetSign(par)
		data = ex.ConnWithHeader("GET", endpoint, par)
		if ex.cfg.DebugMode {
			log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: GET\tEndpoint:%s %v", endpoint, string(data))
		}
	} else {
		data = ex.ConnWithoutHeader(endpoint, par)
	}
	switch endpoint {
	case EndpointAccountInformation:
		var accountInformation response.AccountInformation
		// не может возвращать ошибку
		json.Unmarshal(data, &accountInformation)
		returnInterface = accountInformation
	case EndpointAccountTradeList:
		var accountTradeList response.AccountTradeList
		// не может возвращать ошибку
		json.Unmarshal(data, &accountTradeList)
		returnInterface = accountTradeList
	case EndpointQueryOpenOCO:
		var queryOpenOCO response.QueryOpenOCO
		// не может возвращать ошибку
		json.Unmarshal(data, &queryOpenOCO)
		returnInterface = queryOpenOCO
	case EndpointExchangeInfo:
		var exchangeInformation response.ExchangeInformation
		json.Unmarshal(data, &exchangeInformation)
		returnInterface = exchangeInformation
	default:
		log.Fatal()
	}
	return returnInterface
}

func (ex *BinanceRest) AccountTradeList(parm parameters.AccountTradeList) response.AccountTradeList {

	par := parameters.BinanceParmAccountTradeListToString(parm)
	par += "&signature=" + ex.GetSign(par)
	data := ex.ConnWithHeader("GET", EndpointAccountTradeList, par)
	if ex.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: GET\tEndpoint:%sBinanceWalletBalance %v", EndpointAccountTradeList, string(data))
	}

	var accountTradeList response.AccountTradeList
	if !ex.isErr(data) {
		_ = json.Unmarshal(data, &accountTradeList)
	} else {
		log.Printf(`
				{
					"Status" : "Error",
					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
					"File": "get.go",
					"Functions" : "(ex *BinanceRest) AccountTradeList(parm parameters.AccountTradeList) response.AccountTradeList",
					"Function where err" : "json.Unmarshal",
					"Exchange" : "Binance",
					"Error" : %s
				}`, string(data))
		log.Fatal()
	}

	return accountTradeList
}

func (ex *BinanceRest) QueryOpenOCO(parm parameters.QueryOpenOCO) response.QueryOpenOCO {

	par := ""
	par = parameters.BinanceParmQueryOpenOCOToString(parm)
	par += "&signature=" + ex.GetSign(par)
	data := ex.ConnWithHeader("GET", EndpointQueryOpenOCO, par)
	if ex.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: GET\tEndpoint:%sBinanceWalletBalance %v", EndpointQueryOpenOCO, string(data))
	}
	var queryOpenOCO response.QueryOpenOCO
	if !ex.isErr(data) {
		_ = json.Unmarshal(data, &queryOpenOCO)
	} else {
		log.Printf(`
				{
					"Status" : "Error",
					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
					"File": "get.go",
					"Functions" : "(ex *BinanceRest) QueryOpenOCO(parm parameters.QueryOpenOCO) response.QueryOpenOCO ",
					"Function where err" : "json.Unmarshal",
					"Exchange" : "Binance",
					"Error" : %s
				}`, string(data))
		log.Fatal()
	}
	return queryOpenOCO
}

func (ex *BinanceRest) ExchangeInformation(parm parameters.ExchangeInformation) response.ExchangeInformation {

	par := ""
	par = parameters.BinanceParmExchangeInformationToString(parm)
	data := ex.ConnWithoutHeader(EndpointExchangeInfo, par)
	if ex.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: GET\tEndpoint:%s\tRESPONSE: %v", EndpointExchangeInfo, string(data))
	}
	var exchangeInformation response.ExchangeInformation

	if !ex.isErr(data) {
		_ = json.Unmarshal(data, &exchangeInformation)
	} else {
		log.Printf(`
				{
					"Status" : "Error",
					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
					"File": "get.go",
					"Functions" : "(ex *BinanceRest) ExchangeInformation(parm parameters.ExchangeInformation) response.ExchangeInformation",
					"Function where err" : "json.Unmarshal",
					"Exchange" : "Binance",
					"Error" : %s
				}`, string(data))
		log.Fatal()
	}
	return exchangeInformation
}
