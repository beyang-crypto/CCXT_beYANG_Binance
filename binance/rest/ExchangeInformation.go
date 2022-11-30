package rest

import (
	"log"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#exchange-information
	EndpointExchangeInfo = "/api/v3/exchangeInfo"
)

type ExchangeInformationParam struct {
	Symbol      *string
	Symbols     *[]string
	Permissions *[]string
}

type ExchangeInformationResp struct {
	Timezone        string            `json:"timezone"`
	ServerTime      int64             `json:"serverTime"`
	RateLimits      []RateLimits      `json:"rateLimits"`
	ExchangeFilters []ExchangeFilters `json:"exchangeFilters"`
	Symbols         []Symbol          `json:"symbols"`
}

type Symbol struct {
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
}

// https://binance-docs.github.io/apidocs/spot/en/#public-api-definitions
type RateLimits struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

// https://binance-docs.github.io/apidocs/spot/en/#filters
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

func (ex *BinanceRest) ExchangeInformation(parm ExchangeInformationParam) ExchangeInformationResp {
	r := &Request{
		method:   http.MethodGet,
		endpoint: EndpointExchangeInfo,
		secType:  secTypeNone,
	}

	m := setExchangeInformationParams(parm)
	r.setParams(m)
	data, err := ex.callAPI(r)

	if err != nil {
		log.Printf("%v", err)
	}
	var exchangeInformation ExchangeInformationResp
	_ = json.Unmarshal(data, &exchangeInformation)
	return exchangeInformation
}

func setExchangeInformationParams(parm ExchangeInformationParam) params {
	var m params
	if parm.Symbol != nil {
		m["symbol"] = *parm.Symbol
	}
	if parm.Symbols != nil {
		if len(*parm.Symbols) == 0 {
			m["symbol"] = "[]"
		} else {
			m["symbol"] = "[\"" + strings.Join(*parm.Symbols, "\",\"") + "\"]"
		}
	}
	if parm.Permissions != nil {
		if len(*parm.Permissions) == 0 {
			m["permissions"] = "[]"
		} else {
			m["permissions"] = "[\"" + strings.Join(*parm.Permissions, "\",\"") + "\"]"
		}
	}
	return m
}
