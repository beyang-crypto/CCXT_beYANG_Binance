package ws

const (
	HostMainnetPublicTopics = "wss://stream.binance.com:9443/ws"
)

const (
	// https://binance-docs.github.io/apidocs/spot/en/#kline-candlestick-streams
	ChannelKline = "@kline_"

	// https://binance-docs.github.io/apidocs/spot/en/#individual-symbol-book-ticker-streams
	ChannelTicker = "@bookTicker"
)

const (
	KlineInterval1seconds  string = "1s"
	KlineInterval1minutes  string = "1m"
	KlineInterval3minutes  string = "3m"
	KlineInterval5minutes  string = "5m"
	KlineInterval15minutes string = "15m"
	KlineInterval30minutes string = "30m"
	KlineInterval1hours    string = "1h"
	KlineInterval2hours    string = "2h"
	KlineInterval4hours    string = "4h"
	KlineInterval6hours    string = "6h"
	KlineInterval8hours    string = "8h"
	KlineInterval12hours   string = "12h"
	KlineInterval1days     string = "1d"
	KlineInterval3days     string = "3d"
	KlineInterval1weeks    string = "1w"
	KlineInterval1months   string = "1M"
)
