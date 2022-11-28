package main

import (
	"log"

	config "github.com/beyang-crypto/CCXT_beYANG_Binance/binance"

	binanceRest "github.com/beyang-crypto/CCXT_beYANG_Binance/binance/rest"
)

func main() {

	path := "config-prod.yaml"

	conf := config.GetConfig(path)

	cfg := &binanceRest.Configuration{
		Addr:      binanceRest.BaseEndpoint,
		ApiKey:    conf.Api.Key,
		SecretKey: conf.Api.Secret,
		DebugMode: true,
	}

	b := binanceRest.New(cfg)

	parm := binanceRest.AccountInformationParam{}

	resp := b.AccountInformation(parm)

	log.Printf("response %v", resp)
}
