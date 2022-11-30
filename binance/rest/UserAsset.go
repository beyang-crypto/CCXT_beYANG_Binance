package rest

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#user-asset-user_data
	EndpointUserAsset = "/sapi/v3/asset/getUserAsset"
)

type UserAssetParam struct {
	Asset            *string // optional
	NeedBtcValuation *bool   // optional
	RecvWindow       *int64  // optional
}

type UserAssetResp struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}

func (ex *BinanceRest) UserAsset(parm UserAssetParam) UserAssetResp {
	r := &Request{
		method:   http.MethodPost,
		endpoint: EndpointUserAsset,
		secType:  secTypeAPIKey,
	}

	m := setUserAssetParams(parm)
	r.setParams(m)

	data, err := ex.callAPI(r)
	if err != nil {
		log.Printf("%v", err)
	}

	var userAsset UserAssetResp
	_ = json.Unmarshal(data, &userAsset)
	return userAsset
}

func setUserAssetParams(parm UserAssetParam) params {
	m := params{}
	if parm.Asset != nil {
		m["asset"] = *&parm.Asset
	}
	if parm.NeedBtcValuation != nil {
		m["needBtcValuation"] = *&parm.NeedBtcValuation
	}
	if parm.RecvWindow != nil {
		m["recvWindow"] = *&parm.RecvWindow
	}
	return m
}
