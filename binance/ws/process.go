package ws

func (b *BinanceWS) processBookTicker(symbol string, data BookTicker) {
	b.Emit(ChannelTicker, symbol, data)
}
