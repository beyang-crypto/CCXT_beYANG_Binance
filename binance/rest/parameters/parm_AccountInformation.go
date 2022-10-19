package parameters

import "fmt"

type AccountInformation struct {
	RecvWindow int64 `json:"recvWindow"`
}

func BinanceParmsToAccountInformation(data interface{}) (AccountInformation, bool) {
	bt, ok := data.(AccountInformation)
	return bt, ok
}

func BinanceParmAccountInformationToString(parm AccountInformation) string {
	par := ""
	par += checkRecvWindow(parm.RecvWindow)
	par += fmt.Sprintf("timestamp=%d", getTimestamp())
	return par
}
