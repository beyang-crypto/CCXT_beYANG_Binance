package rest

// func (ex *BinanceRest) AccountTradeList(parm parameters.AccountTradeList) response.AccountTradeList {

// 	par := parameters.BinanceParmAccountTradeListToString(parm)
// 	par += "&signature=" + ex.GetSign(par)
// 	data := ex.ConnWithHeader("GET", EndpointAccountTradeList, par)
// 	if ex.cfg.DebugMode {
// 		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: GET\tEndpoint:%sBinanceWalletBalance %v", EndpointAccountTradeList, string(data))
// 	}

// 	var accountTradeList response.AccountTradeList
// 	if !ex.isErr(data) {
// 		_ = json.Unmarshal(data, &accountTradeList)
// 	} else {
// 		log.Printf(`
// 				{
// 					"Status" : "Error",
// 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
// 					"File": "get.go",
// 					"Functions" : "(ex *BinanceRest) AccountTradeList(parm parameters.AccountTradeList) response.AccountTradeList",
// 					"Function where err" : "json.Unmarshal",
// 					"Exchange" : "Binance",
// 					"Error" : %s
// 				}`, string(data))
// 		log.Fatal()
// 	}

// 	return accountTradeList
// }

// func (ex *BinanceRest) QueryOpenOCO(parm parameters.QueryOpenOCO) response.QueryOpenOCO {

// 	par := ""
// 	par = parameters.BinanceParmQueryOpenOCOToString(parm)
// 	par += "&signature=" + ex.GetSign(par)
// 	data := ex.ConnWithHeader("GET", EndpointQueryOpenOCO, par)
// 	if ex.cfg.DebugMode {
// 		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: GET\tEndpoint:%sBinanceWalletBalance %v", EndpointQueryOpenOCO, string(data))
// 	}
// 	var queryOpenOCO response.QueryOpenOCO
// 	if !ex.isErr(data) {
// 		_ = json.Unmarshal(data, &queryOpenOCO)
// 	} else {
// 		log.Printf(`
// 				{
// 					"Status" : "Error",
// 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
// 					"File": "get.go",
// 					"Functions" : "(ex *BinanceRest) QueryOpenOCO(parm parameters.QueryOpenOCO) response.QueryOpenOCO ",
// 					"Function where err" : "json.Unmarshal",
// 					"Exchange" : "Binance",
// 					"Error" : %s
// 				}`, string(data))
// 		log.Fatal()
// 	}
// 	return queryOpenOCO
// }

// func (ex *BinanceRest) ExchangeInformation(parm parameters.ExchangeInformation) response.ExchangeInformation {

// 	par := ""
// 	par = parameters.BinanceParmExchangeInformationToString(parm)
// 	data := ex.ConnWithoutHeader("GET", EndpointExchangeInfo, par)
// 	if ex.cfg.DebugMode {
// 		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: GET\tEndpoint:%s\tRESPONSE: %v", EndpointExchangeInfo, string(data))
// 	}
// 	var exchangeInformation response.ExchangeInformation

// 	if !ex.isErr(data) {
// 		_ = json.Unmarshal(data, &exchangeInformation)
// 	} else {
// 		log.Printf(`
// 				{
// 					"Status" : "Error",
// 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
// 					"File": "get.go",
// 					"Functions" : "(ex *BinanceRest) ExchangeInformation(parm parameters.ExchangeInformation) response.ExchangeInformation",
// 					"Function where err" : "json.Unmarshal",
// 					"Exchange" : "Binance",
// 					"Error" : %s
// 				}`, string(data))
// 		log.Fatal()
// 	}
// 	return exchangeInformation
// }
