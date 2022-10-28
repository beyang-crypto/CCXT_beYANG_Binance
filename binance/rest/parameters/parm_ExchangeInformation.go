package parameters

import (
	"log"
)

type ExchangeInformation struct {
	Symbols     []string
	Permissions []string
}

func BinanceToExchangeInformationParms(data interface{}) (ExchangeInformation, bool) {
	ei, ok := data.(ExchangeInformation)
	return ei, ok
}

func BinanceParmExchangeInformationToString(parm ExchangeInformation) string {
	strParm := ""
	switch len(parm.Symbols) {
	case 0:
		if len(parm.Permissions) != 0 {
			strParm = checkPermissions(parm.Permissions...)
		}
	case 1:
		strParm = checkSymbol(parm.Symbols[0])
	default:
		strParm = checkSymbols(parm.Symbols...)
	}
	log.Printf(strParm)
	return strParm
}
