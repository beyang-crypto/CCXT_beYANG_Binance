package main

import (
	"log"

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

	pair1 := b.GetPair("btc", "usdt")
	pair2 := b.GetPair("eth", "usdt")
	pair3 := b.GetPair("xrp", "usdt")
	pair4 := b.GetPair("ada", "usdt")
	pair5 := b.GetPair("sol", "usdt")
	pair6 := b.GetPair("doge", "usdt")
	pair7 := b.GetPair("matic", "usdt")
	pair8 := b.GetPair("shib", "usdt")
	pair9 := b.GetPair("trx", "usdt")
	pair10 := b.GetPair("uni", "usdt")
	pair11 := b.GetPair("avax", "usdt")
	pair12 := b.GetPair("ltc", "usdt")
	pair13 := b.GetPair("etc", "usdt")
	pair14 := b.GetPair("link", "usdt")
	pair15 := b.GetPair("atom", "usdt")

	b.Subscribe(binanceWs.ChannelTicker, []string{pair1, pair2, pair3, pair4, pair5, pair6, pair7, pair8, pair9, pair10, pair11, pair12, pair13, pair14, pair15})

	b.On(binanceWs.ChannelTicker, handleBookTicker)
	b.On(binanceWs.ChannelTicker, handleBestBidPrice)

	// cfgRest := &binanceRest.Configuration{
	// 	Addr:      binanceRest.RestBaseEndpoint,
	// 	ApiKey:    "",
	// 	SecretKey: "",
	// 	DebugMode: true,
	// }
	// r := binanceRest.New(cfgRest)
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	balance := binanceRest.BinanceToWalletBalance(r.GetBalance())
	// 	for _, coins := range balance.Balances {
	// 		log.Printf("coin = %s, total = %s", coins.Asset, coins.Free)
	// 	}
	// }()

	//	не дает прекратить работу программы
	forever := make(chan struct{})
	<-forever
}

func handleBookTicker(name string, symbol string, data binanceWs.BookTicker) {
	log.Printf("%s Ticker  %s: %v", name, symbol, data)
}

func handleBestBidPrice(name string, symbol string, data binanceWs.BookTicker) {
	log.Printf("%s BookTicker  %s: BestBidPrice : %s", name, symbol, data.B)
}
