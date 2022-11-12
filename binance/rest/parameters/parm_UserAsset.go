package parameters

import (
	"fmt"
)

// https://binance-docs.github.io/apidocs/spot/en/#user-asset-user_data
type UserAsset struct {
	Asset            string // optional
	NeedBtcValuation bool   // optional
	RecvWindow       int64  // optional
}

func BinanceParmsToUserAsset(data interface{}) (UserAsset, bool) {
	ua, ok := data.(UserAsset)
	return ua, ok
}

func BinanceParmUserAssetToString(parm UserAsset) string {
	par := ""
	par += checkAsset(parm.Asset)
	par += checkNeedBtcValuation(parm.NeedBtcValuation)
	par += checkRecvWindow(parm.RecvWindow)
	par += fmt.Sprintf("timestamp=%d", getTimestamp())
	return par
}
