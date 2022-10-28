package response

type ExchangeInformation struct {
	Timezone        string            `json:"timezone"`
	ServerTime      int64             `json:"serverTime"`
	RateLimits      []RateLimits      `json:"rateLimits"`
	ExchangeFilters []ExchangeFilters `json:"exchangeFilters"`
	Symbols         []struct {
		Symbol                     string          `json:"symbol"`
		Status                     string          `json:"status"`
		BaseAsset                  string          `json:"baseAsset"`
		BaseAssetPrecision         int             `json:"baseAssetPrecision"`
		QuoteAsset                 string          `json:"quoteAsset"`
		QuotePrecision             int             `json:"quotePrecision"`
		QuoteAssetPrecision        int             `json:"quoteAssetPrecision"`
		BaseCommissionPrecision    int             `json:"baseCommissionPrecision"`
		QuoteCommissionPrecision   int             `json:"quoteCommissionPrecision"`
		OrderTypes                 []string        `json:"orderTypes"`
		IcebergAllowed             bool            `json:"icebergAllowed"`
		OcoAllowed                 bool            `json:"ocoAllowed"`
		QuoteOrderQtyMarketAllowed bool            `json:"quoteOrderQtyMarketAllowed"`
		AllowTrailingStop          bool            `json:"allowTrailingStop"`
		CancelReplaceAllowed       bool            `json:"cancelReplaceAllowed"`
		IsSpotTradingAllowed       bool            `json:"isSpotTradingAllowed"`
		IsMarginTradingAllowed     bool            `json:"isMarginTradingAllowed"`
		Filters                    []SymbolFilters `json:"filters"`
		Permissions                []string        `json:"permissions"`
	} `json:"symbols"`
}

func BinanceToExchangeInformation(data interface{}) (ExchangeInformation, bool) {
	ei, ok := data.(ExchangeInformation)
	return ei, ok
}
