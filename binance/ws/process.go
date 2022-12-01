package ws

func (b *BinanceWS) processBookTicker(name string, symbol string, data BookTicker) {
	b.Emit(ChannelTicker, name, symbol, data)
}

func (b *BinanceWS) processKline(name string, symbol string, data Kline) {
	b.Emit(ChannelTicker, name, symbol, data)
}
