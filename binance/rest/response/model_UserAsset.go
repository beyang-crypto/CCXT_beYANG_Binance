package response

type UserAsset struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}

func BinanceToUserAsser(data interface{}) (UserAsset, bool) {
	ua, ok := data.(UserAsset)
	return ua, ok
}
