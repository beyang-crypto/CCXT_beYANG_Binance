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
	symbol := b.GetPair("btc", "usdt")
	parm := binanceRest.MarginAccountPriceIndexParam{
		Symbol: symbol,
	}

	resp := b.MarginAccountPriceIndex(parm)

	log.Printf("response %v", resp)
}
