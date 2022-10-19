package response

type TestNewOrder struct {
}

func BinanceToTestNewOrder(data interface{}) (TestNewOrder, bool) {
	tno, ok := data.(TestNewOrder)
	return tno, ok
}
