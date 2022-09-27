package main

import (
	"log"
	"time"

	binanceRest "github.com/TestingAccMar/CCXT_beYANG_Binance/binance/rest"
	binanceWs "github.com/TestingAccMar/CCXT_beYANG_Binance/binance/ws"
)

func main() {
	cfg := &binanceWs.Configuration{
		Addr:      binanceWs.HostMainnetPublicTopics,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: true,
	}
	b := binanceWs.New(cfg)
	b.Start()

	//pair1 := b.GetPair("btc", "usdt")

	//b.Subscribe(binanceWs.ChannelTicker, pair1)
	//b.Subscribe(binanceWs.ChannelTicker, "ETHBTC")

	// b.On(binanceWs.ChannelTicker, handleBookTicker)
	// b.On(binanceWs.ChannelTicker, handleBestBidPrice)

	cfgRest := &binanceRest.Configuration{
		Addr:      binanceRest.RestBaseEndpoint,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: true,
	}
	r := binanceRest.New(cfgRest)
	go func() {
		time.Sleep(1 * time.Second)
		balance := binanceRest.BinanceToWalletBalance(r.GetBalance())
		for _, coins := range balance.Balances {
			log.Printf("coin = %s, total = %s", coins.Asset, coins.Free)
		}
	}()

	//	не дает прекратить работу программы
	forever := make(chan struct{})
	<-forever
}

func handleBookTicker(symbol string, data binanceWs.BookTicker) {
	log.Printf("Binance Ticker  %s: %v", symbol, data)
}

func handleBestBidPrice(symbol string, data binanceWs.BookTicker) {
	log.Printf("Binance BookTicker  %s: BestBidPrice : %s", symbol, data.B)
}
