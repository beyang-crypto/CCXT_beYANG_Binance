package v3

func (b *BinanceWS) processBookTicker(symbol string, data BookTicker) {
	b.Emit(ChannelTicker, symbol, data)
}
