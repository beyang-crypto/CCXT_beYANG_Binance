package main

import (
	"log"
	"time"

	binanceRest "github.com/TestingAccMar/CCXT_beYANG_Binance/binance/rest"
	binanceRestParm "github.com/TestingAccMar/CCXT_beYANG_Binance/binance/rest/parameters"
	"github.com/TestingAccMar/CCXT_beYANG_Binance/binance/rest/response"
)

func main() {
	cfg := &binanceRest.Configuration{
		Addr:      binanceRest.BaseEndpoint,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: true,
	}

	b := binanceRest.New(cfg)
	parm1 := binanceRestParm.ExchangeInformation{
		Permissions: []string{"MARGIN", "LEVERAGED"},
	}
	a := b.ExchangeInformation(parm1)

	log.Printf("%v", a)

	time.Sleep(1 * time.Second)
	log.Printf("========================================================================================================================================================================")
	time.Sleep(1 * time.Second)

	parm := binanceRestParm.AccountInformation{}

	i := b.Get(binanceRest.EndpointAccountInformation, parm)
	i1, ok := response.BinanceToAccountInformation(i)
	if ok {
		log.Printf("%v", i1)
	}
}
