package bitflyer

// ####################
// Product Code
// ####################
type ProductCode int

const (
	Btcjpy ProductCode = iota
	Ethjpy
	Fxbtcjpy
	Ethbtc
	bchbtc
)

func (code ProductCode) String() string {
	switch code {
	case Btcjpy:
		return "BTC_JPY"
	case Ethjpy:
		return "ETH_JPY"
	case Fxbtcjpy:
		return "FX_BTC_JPY"
	case Ethbtc:
		return "ETH_BTC"
	case bchbtc:
		return "BCH_BYC"
	default:
		return "BTC_JPY"
	}
}

// ####################
// Order Type
// ####################

type OrderType int

const (
	Limit OrderType = iota
	Market
)

func (ot OrderType) String() string {
	switch ot {
	case Limit:
		return "LIMIT"
	case Market:
		return "MARKET"
	default:
		return "LIMIT"
	}
}

// ####################
// Side
// ####################
type Side int

const (
	Buy Side = iota
	Sell
)

func (ot Side) String() string {
	switch ot {
	case Buy:
		return "BUY"
	case Sell:
		return "Sell"
	default:
		return "BUY"
	}
}

// ####################
// Time In Force
// ####################
type TimeInForce int

const (
	Gtc TimeInForce = iota
	Ioc
	Fok
)

func (ot TimeInForce) String() string {
	switch ot {
	case Gtc:
		return "GTC"
	case Ioc:
		return "IOC"
	case Fok:
		return "FOK"
	default:
		return "GTC"
	}
}
