package ws

// https://binance-docs.github.io/apidocs/spot/en/#kline-candlestick-streams
type Kline struct {
	E  string `json:"e"`
	E0 int    `json:"E"`
	S  string `json:"s"`
	K  struct {
		T  int    `json:"t"`
		T0 int    `json:"T"`
		S  string `json:"s"`
		I  string `json:"i"`
		F  int    `json:"f"`
		L  int    `json:"L"`
		O  string `json:"o"`
		C  string `json:"c"`
		H  string `json:"h"`
		L0 string `json:"l"`
		V  string `json:"v"`
		N  int    `json:"n"`
		X  bool   `json:"x"`
		Q  string `json:"q"`
		V0 string `json:"V"`
		Q0 string `json:"Q"`
		B  string `json:"B"`
	} `json:"k"`
}

// https://binance-docs.github.io/apidocs/spot/en/#individual-symbol-book-ticker-streams
type BookTicker struct {
	U  int64  `json:"u"`
	S  string `json:"s"`
	B  string `json:"b"`
	B0 string `json:"B"`
	A  string `json:"a"`
	A0 string `json:"A"`
}
