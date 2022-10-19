package parameters

import "fmt"

//	string

func checkSymbol(s string) string {
	if s != "" {
		return fmt.Sprintf("symbol=%s&", s)
	} else {
		return ""
	}
}

func checkSide(s string) string {
	if s != "" {
		return fmt.Sprintf("side=%s&", s)
	} else {
		return ""
	}
}

func checkType(s string) string {
	if s != "" {
		return fmt.Sprintf("type=%s&", s)
	} else {
		return ""
	}
}

func checkTimeInForce(s string) string {
	if s != "" {
		return fmt.Sprintf("timeInForce=%s&", s)
	} else {
		return ""
	}
}

func checkOrigClientOrderId(s string) string {
	if s != "" {
		return fmt.Sprintf("origClientOrderId=%s&", s)
	} else {
		return ""
	}
}

// int64

func checkOrderId(oid int64) string {
	if oid != 0 {
		return fmt.Sprintf("orderId=%d&", oid)
	} else {
		return ""
	}
}

func checkStartTime(st int64) string {
	if st != 0 {
		return fmt.Sprintf("startTime=%d&", st)
	} else {
		return ""
	}
}

func checkEndTime(et int64) string {
	if et != 0 {
		return fmt.Sprintf("endTime=%d&", et)
	} else {
		return ""
	}
}

func checkFromId(fi int64) string {
	if fi != 0 {
		return fmt.Sprintf("fromId=%d&", fi)
	} else {
		return ""
	}
}

func checkLimit(l int64) string {
	if l != 0 {
		return fmt.Sprintf("limit=%d&", l)
	} else {
		return ""
	}
}

func checkRecvWindow(r int64) string {
	if r != 0 {
		return fmt.Sprintf("recvWindow=%d&", r)
	} else {
		return ""
	}
}

// float64
func checkQuantity(r float64) string {
	if r != 0 {
		return fmt.Sprintf("quantity=%f&", r)
	} else {
		return ""
	}
}

func checkQuoteOrderQty(r float64) string {
	if r != 0 {
		return fmt.Sprintf("quoteOrderQty=%f&", r)
	} else {
		return ""
	}
}

func checkPrice(r float64) string {
	if r != 0 {
		return fmt.Sprintf("price=%f&", r)
	} else {
		return ""
	}
}
