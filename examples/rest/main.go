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
		ApiKey:    "sJsDdkKA0MrQSethzPEyUTzwL6Ui19CtiqaBGbT8PmTcl5EDeMiY6gMqrANYGADz",
		SecretKey: "3y1wB3mjJ0Rg8BuAor3nP8SO02oMBeyof6CqotVD2Yd9ARyWaercTBzOOkRM9p0b",
		DebugMode: false,
	}

	b := binanceRest.New(cfg)
	pair := b.GetPair("btc", "usdt")
	parm := binanceRestParm.TestNewOrder{
		Symbol:      pair,
		Side:        "BUY",
		Type:        "LIMIT",
		Quantity:    1,
		TimeInForce: "GTC",
		Price:       19158.06,
		//RecvWindow: 60000,
	}
	b.TestNewOrder(parm)

	time.Sleep(1 * time.Second)
	log.Printf("=================================================================================================================================================================================================================================")
	time.Sleep(1 * time.Second)

	parm = binanceRestParm.TestNewOrder{
		Symbol:      pair,
		Side:        "BUY",
		Type:        "LIMIT",
		Quantity:    1,
		TimeInForce: "GTC",
		Price:       19158.06,
		RecvWindow:  60000,
	}

	i := b.Post(binanceRest.EndpointTestNewOrder, parm)
	i1, ok := response.BinanceToTestNewOrder(i)
	if ok {
		log.Printf("%v", i1)
	}
}
