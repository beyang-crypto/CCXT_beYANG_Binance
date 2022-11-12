package rest

import (
	"log"

	"github.com/beyang-crypto/CCXT_beYANG_Binance/binance/rest/parameters"
	"github.com/beyang-crypto/CCXT_beYANG_Binance/binance/rest/response"
	"github.com/goccy/go-json"
)

func (ex *BinanceRest) Post(endpoint string, parms interface{}) interface{} {

	parmsForEp := false
	withAuth := false
	par := ""
	switch endpoint {
	case EndpointTestNewOrder:
		parm, ok := parameters.BinanceParmsToTestNewOrder(parms)
		if ok {
			par = parameters.BinanceParmTestNewOrderToString(parm)
			withAuth = true
		} else {
			parmsForEp = true
		}
	case EndpointNewOrder:
		parm, ok := parameters.BinanceParmsToTestNewOrder(parms)
		if ok {
			par = parameters.BinanceParmTestNewOrderToString(parm)
			withAuth = true
		} else {
			parmsForEp = true
		}
	case EndpointUserAsset:
		parm, ok := parameters.BinanceParmsToUserAsset(parms)
		if ok {
			par = parameters.BinanceParmUserAssetToString(parm)
			withAuth = true
		} else {
			parmsForEp = true
		}
	default:
		log.Printf(`
			{
				"Status" : "Error",
				"Path to file" : "CCXT_beYANG_Binance/binance/rest",
				"File": "post.go",
				"Functions" : "(ex *BinanceRest) Post(endpoint string, parms interface{})",
				"Exchange" : "Binance",
				"Error" : Unable to connect to this endpoint,
			}`)
		log.Fatal()
	}
	// Если параметр не соответствует конечной точке
	if parmsForEp {
		log.Printf(`
			{
				"Status" : "Error",
				"Path to file" : "CCXT_beYANG_Binance/binance/rest",
				"File": "post.go",
				"Functions" : "(ex *BinanceRest) Post(endpoint string, parms interface{})",
				"Exchange" : "Binance",
				"Error" : the parameters do not match the endpoint(%s),
			}`, endpoint)
		log.Fatal()
	}

	var data []byte

	var returnInterface interface{}
	if withAuth {
		par += "&signature=" + ex.GetSign(par)
		data = ex.ConnWithHeader("POST", endpoint, par)
		if ex.cfg.DebugMode {
			log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: POST\tEndpoint:%s %v", endpoint, string(data))
		}
	}
	switch endpoint {
	case EndpointTestNewOrder:
		var testNewOrder response.TestNewOrder
		// не может возвращать ошибку
		json.Unmarshal(data, testNewOrder)
		returnInterface = testNewOrder
	default:
		log.Fatal()
	}
	return returnInterface
}

func (ex *BinanceRest) TestNewOrder(parm parameters.TestNewOrder) response.TestNewOrder {
	par := ""
	par = parameters.BinanceParmTestNewOrderToString(parm)
	par += "&signature=" + ex.GetSign(par)
	data := ex.ConnWithHeader("POST", EndpointTestNewOrder, par)
	var testNewOrder response.TestNewOrder
	if ex.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: POST\tEndpoint:%s TestNewOrder  %v", EndpointTestNewOrder, string(data))
	}
	if !ex.isErr(data) {
		_ = json.Unmarshal(data, &testNewOrder)

	} else {
		log.Printf(`
				{
					"Status" : "Error",
					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
					"File": "post.go",
					"Functions" : "(ex *BinanceRest) GetTestNewOrder(parm parameters.TestNewOrder) response.TestNewOrder",
					"Function where err" : "json.Unmarshal",
					"Exchange" : "Binance",
					"Error" : %s
				}`, string(data))
		log.Fatal()
	}
	return testNewOrder
}

func (ex *BinanceRest) NewOrder(parm parameters.NewOrder) response.NewOrder {
	par := ""
	par = parameters.BinanceParmNewOrderToString(parm)
	par += "&signature=" + ex.GetSign(par)
	data := ex.ConnWithHeader("POST", EndpointNewOrder, par)
	var newOrder response.NewOrder
	if ex.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: POST\tEndpoint:%s NewOrder  %v", EndpointNewOrder, string(data))
	}
	if !ex.isErr(data) {
		_ = json.Unmarshal(data, &newOrder)
	} else {
		log.Printf(`
				{
					"Status" : "Error",
					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
					"File": "post.go",
					"Functions" : "(ex *BinanceRest) NewOrder(parm parameters.NewOrder) response.NewOrder",
					"Function where err" : "json.Unmarshal",
					"Exchange" : "Binance",
					"Error" : %s
				}`, string(data))
		log.Fatal()
	}
	return newOrder
}

func (ex *BinanceRest) UserAsset(parm parameters.UserAsset) response.UserAsset {
	par := ""
	par = parameters.BinanceParmUserAssetToString(parm)
	par += "&signature=" + ex.GetSign(par)
	data := ex.ConnWithHeader("POST", EndpointUserAsset, par)
	var userAsset response.UserAsset
	if ex.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tMethod: POST\tEndpoint:%s UserAsset  %v", EndpointUserAsset, string(data))
	}
	if !ex.isErr(data) {
		_ = json.Unmarshal(data, &userAsset)
	} else {
		log.Printf(`
				{
					"Status" : "Error",
					"Path to file" : "CCXT_beYANG_Binance/binance/rest",
					"File": "post.go",
					"Functions" : "(ex *BinanceRest) UserAsset(parm parameters.UserAsset) response.UserAsset",
					"Function where err" : "json.Unmarshal",
					"Exchange" : "Binance",
					"Error" : %s
				}`, string(data))
		log.Fatal()
	}
	return userAsset
}
