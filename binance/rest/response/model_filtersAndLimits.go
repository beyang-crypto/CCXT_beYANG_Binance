package response

type RateLimits struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

type SymbolFilters struct {
	FilterType            string `json:"filterType"`
	MinPrice              string `json:"minPrice"`
	MaxPrice              string `json:"maxPrice"`
	TickSize              string `json:"tickSize"`
	MultiplierUp          string `json:"multiplierUp"`
	MultiplierDown        string `json:"multiplierDown"`
	AvgPriceMins          int    `json:"avgPriceMins"`
	BidMultiplierUp       string `json:"bidMultiplierUp"`
	BidMultiplierDown     string `json:"bidMultiplierDown"`
	AskMultiplierUp       string `json:"askMultiplierUp"`
	AskMultiplierDown     string `json:"askMultiplierDown"`
	MinQty                string `json:"minQty"`
	MaxQty                string `json:"maxQty"`
	StepSize              string `json:"stepSize"`
	MinNotional           string `json:"minNotional"`
	ApplyMinToMarket      bool   `json:"applyMinToMarket"`
	MaxNotional           string `json:"maxNotional"`
	ApplyMaxToMarket      bool   `json:"applyMaxToMarket"`
	ApplyToMarket         bool   `json:"applyToMarket"`
	Limit                 int    `json:"limit"`
	MinTrailingAboveDelta int    `json:"minTrailingAboveDelta"`
	MaxTrailingAboveDelta int    `json:"maxTrailingAboveDelta"`
	MinTrailingBelowDelta int    `json:"minTrailingBelowDelta"`
	MaxTrailingBelowDelta int    `json:"maxTrailingBelowDelta"`
	MaxNumOrders          int    `json:"maxNumOrders"`
	MaxNumAlgoOrders      int    `json:"maxNumAlgoOrders"`
	MaxNumIcebergOrders   int    `json:"maxNumIcebergOrders"`
	MaxPosition           string `json:"maxPosition"`
}

type ExchangeFilters struct {
	FilterType          string `json:"filterType"`
	MaxNumOrders        int    `json:"maxNumOrders"`
	MaxNumAlgoOrders    int    `json:"maxNumAlgoOrders"`
	MaxNumIcebergOrders int    `json:"maxNumIcebergOrders"`
}
