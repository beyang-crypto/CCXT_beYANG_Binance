package rest

//Endpoint
const (
	//https://binance-docs.github.io/apidocs/spot/en/#test-new-order-trade
	EndpointTestNewOrder = "/api/v3/order/test"
	//https://binance-docs.github.io/apidocs/spot/en/#new-order-trade
	EndpointNewOrder = "/api/v3/order"
	//https://binance-docs.github.io/apidocs/spot/en/#query-order-user_data
	EndpointQueryOrder = "/api/v3/order"
	//https://binance-docs.github.io/apidocs/spot/en/#query-open-oco-user_data
	EndpointQueryOpenOCO = "/api/v3/openOrderList"
	//https://binance-docs.github.io/apidocs/spot/en/#account-information-user_data
	EndpointAccountInformation = "/api/v3/account"
	//https://binance-docs.github.io/apidocs/spot/en/#account-trade-list-user_data
	EndpointAccountTradeList = "/api/v3/myTrades"
)
