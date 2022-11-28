package rest

// func (ex *BinanceRest) NewOrder(parm parameters.NewOrder) response.NewOrder {
// 	par := ""
// 	par = parameters.BinanceParmNewOrderToString(parm)
// 	par += "&signature=" + ex.GetSign(par)
// 	data := ex.ConnWithHeader("POST", EndpointNewOrder, par)
// 	var newOrder response.NewOrder
// 	if ex.cfg.DebugMode {
// 		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: POST\tEndpoint:%s NewOrder  %v", EndpointNewOrder, string(data))
// 	}
// 	if !ex.isErr(data) {
// 		_ = json.Unmarshal(data, &newOrder)
// 	} else {
// 		log.Printf(`
// 				{
// 					"Status" : "Error",
// 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
// 					"File": "post.go",
// 					"Functions" : "(ex *BinanceRest) NewOrder(parm parameters.NewOrder) response.NewOrder",
// 					"Function where err" : "json.Unmarshal",
// 					"Exchange" : "Binance",
// 					"Error" : %s
// 				}`, string(data))
// 		log.Fatal()
// 	}
// 	return newOrder
// }

// func (ex *BinanceRest) UserAsset(parm parameters.UserAsset) response.UserAsset {
// 	par := ""
// 	par = parameters.BinanceParmUserAssetToString(parm)
// 	par += "&signature=" + ex.GetSign(par)
// 	data := ex.ConnWithHeader("POST", EndpointUserAsset, par)
// 	var userAsset response.UserAsset
// 	if ex.cfg.DebugMode {
// 		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: POST\tEndpoint:%s UserAsset  %v", EndpointUserAsset, string(data))
// 	}
// 	if !ex.isErr(data) {
// 		_ = json.Unmarshal(data, &userAsset)
// 	} else {
// 		log.Printf(`
// 				{
// 					"Status" : "Error",
// 					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
// 					"File": "post.go",
// 					"Functions" : "(ex *BinanceRest) UserAsset(parm parameters.UserAsset) response.UserAsset",
// 					"Function where err" : "json.Unmarshal",
// 					"Exchange" : "Binance",
// 					"Error" : %s
// 				}`, string(data))
// 		log.Fatal()
// 	}
// 	return userAsset
// }
