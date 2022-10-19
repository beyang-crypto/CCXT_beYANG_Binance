package parameters

import "fmt"

type QueryOpenOCO struct {
	RecvWindow int64 `json:"recvWindow"`
}

func BinanceParmsToQueryOpenOCO(data interface{}) (QueryOpenOCO, bool) {
	qooco, ok := data.(QueryOpenOCO)
	return qooco, ok
}

func BinanceParmQueryOpenOCOToString(parm QueryOpenOCO) string {
	par := ""
	par += checkRecvWindow(parm.RecvWindow)
	par += fmt.Sprintf("timestamp=%d", getTimestamp())
	return par
}
