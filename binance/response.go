package v3

type BookTicker struct {
	U  int    `json:"u"`
	S  string `json:"s"`
	B  string `json:"b"`
	B0 string `json:"B"`
	A  string `json:"a"`
	A0 string `json:"A"`
}

type WalletBalance struct {
	MakerCommission  int       `json:"makerCommission"`
	TakerCommission  int       `json:"takerCommission"`
	BuyerCommission  int       `json:"buyerCommission"`
	SellerCommission int       `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	Brokered         bool      `json:"brokered"`
	UpdateTime       int       `json:"updateTime"`
	AccountType      string    `json:"accountType"`
	Balances         []balance `json:"balances"`
	Permissions      []string  `json:"permissions"`
}
type balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}
