package rest

// https://binance-docs.github.io/apidocs/spot/en/#public-api-definitions

type SymbolStatusType string

type OrderType string

type OrderSideType string

// Description see https://binance-docs.github.io/apidocs/spot/en/#new-order-trade
type NewOrderRespType string

// This sets how long an order will be active before expiration
type TimeInForceType string

// Description see https://binance-docs.github.io/apidocs/spot/en/#margin-account-new-order-trade
type SideEffectType string

const (
	SymbolStatusPreTrading   SymbolStatusType = "PRE_TRADING"
	SymbolStatusTrading      SymbolStatusType = "TRADING"
	SymbolStatusPostTrading  SymbolStatusType = "POST_TRADING"
	SymbolStatusEndOfDay     SymbolStatusType = "END_OF_DAY"
	SymbolStatusHalt         SymbolStatusType = "HALT"
	SymbolStatusAuctionMatch SymbolStatusType = "AUCTION_MATCH"
	SymbolStatusBreak        SymbolStatusType = "BREAK"

	AccountPermissionsSpot      string = "SPOT"
	AccountPermissionsMargin    string = "MARGIN"
	AccountPermissionsLeverged  string = "LEVERAGED"
	AccountPermissionsTrdGpr002 string = "TRD_GRP_002"
	AccountPermissionsTrdGpr003 string = "TRD_GRP_003"
	AccountPermissionsTrdGpr004 string = "TRD_GRP_004"
	AccountPermissionsTrdGpr005 string = "TRD_GRP_005"

	SymbolPermissionsSPOT      string = "SPOT"
	SymbolPermissionsMARGIN    string = "MARGIN"
	SymbolPermissionsLEVERAGED string = "LEVERAGED"
	SymbolPermissionsTrdGpr002 string = "TRD_GRP_002"
	SymbolPermissionsTrdGpr003 string = "TRD_GRP_003"
	SymbolPermissionsTrdGpr004 string = "TRD_GRP_004"
	SymbolPermissionsTrdGpr005 string = "TRD_GRP_005"

	OrderTypeLimit           OrderType = "LIMIT"
	OrderTypeMarket          OrderType = "MARKET"
	OrderTypeStopLoss        OrderType = "STOP_LOSS"
	OrderTypeStopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	OrderTypeTakeProfit      OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"
	OrderTypeLimitMaker      OrderType = "LIMIT_MAKER"

	OrderSideBUY  OrderSideType = "BUY"
	OrderSideSELL OrderSideType = "SELL"

	// Good Til Canceled
	// An order will be on the book unless the order is canceled
	TimeInForceGTC TimeInForceType = "GTC"

	// Immediate Or Cancel
	// An order will try to fill the order as much as it can before the order expires
	TimeInForceIOC TimeInForceType = "IOC"

	// Fill or Kill
	// An order will expire if the full order cannot be filled upon execution.
	TimeInForceFOK TimeInForceType = "FOK"

	NewOrderRespACK    NewOrderRespType = "ACK"
	NewOrderRespRESULT NewOrderRespType = "RESULT"
	NewOrderRespFULL   NewOrderRespType = "FULL"

	SideEffectTypeNoSideEffect SideEffectType = "NO_SIDE_EFFECT"
	SideEffectTypeMarginBuy    SideEffectType = "MARGIN_BUY"
	SideEffectTypeAutoRepay    SideEffectType = "AUTO_REPAY"
)
