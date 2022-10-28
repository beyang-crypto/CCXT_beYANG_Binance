package response

type AccountInformation struct {
	MakerCommission  int
	TakerCommission  int
	BuyerCommission  int
	SellerCommission int
	CanTrade         bool
	CanWithdraw      bool
	CanDeposit       bool
	Brokered         bool
	UpdateTime       int
	AccountType      string
	Balances         []Balance
	Permissions      []string
}
type Balance struct {
	Asset  string
	Free   string
	Locked string
}

func BinanceToAccountInformation(data interface{}) (AccountInformation, bool) {
	bt, ok := data.(AccountInformation)
	return bt, ok
}
