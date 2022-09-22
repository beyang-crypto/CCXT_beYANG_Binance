package main

import (
	"log"

	binanceSpotV3 "github.com/TestingAccMar/CCXT_beYANG_Binance/binance"
)

func main() {
	cfg := &binanceSpotV3.Configuration{
		Addr:      binanceSpotV3.HostMainnetPublicTopics,
		ApiKey:    "",
		SecretKey: "",
		DebugMode: true,
	}
	b := binanceSpotV3.New(cfg)
	b.Start()

	pair1 := b.GetPair("btc", "usdt")

	b.Subscribe(binanceSpotV3.ChannelTicker, pair1)
	//b.Subscribe(binanceSpotV3.ChannelTicker, "ETHBTC")

	// b.On(binanceSpotV3.ChannelTicker, handleBookTicker)
	// b.On(binanceSpotV3.ChannelTicker, handleBestBidPrice)

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	balance := b.GetBalance()
	// 	for _, coins := range balance.Balances {
	// 		log.Printf("coin = %s, total = %s", coins.Asset, coins.Free)
	// 	}
	// }()

	//	не дает прекратить работу программы
	forever := make(chan struct{})
	<-forever
}

func handleBookTicker(symbol string, data binanceSpotV3.BookTicker) {
	log.Printf("Binance Ticker  %s: %v", symbol, data)
}

func handleBestBidPrice(symbol string, data binanceSpotV3.BookTicker) {
	log.Printf("Binance BookTicker  %s: BestBidPrice : %s", symbol, data.B)
}
