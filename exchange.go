package taapi

// Exchange represents a cryptocurrency exchange supported by taapi.io
type Exchange string

const (
	ExchangeBinance       Exchange = "binance"
	ExchangeBinanceUS     Exchange = "binanceus"
	ExchangeBinanceUSDM   Exchange = "binanceusdm"
	ExchangeBitfinex      Exchange = "bitfinex"
	ExchangeBitget        Exchange = "bitget"
	ExchangeBitmex        Exchange = "bitmex"
	ExchangeBitstamp      Exchange = "bitstamp"
	ExchangeBybit         Exchange = "bybit"
	ExchangeCoinbase      Exchange = "coinbase"
	ExchangeCryptoCom     Exchange = "cryptocom"
	ExchangeGateIO        Exchange = "gateio"
	ExchangeHuobi         Exchange = "huobi"
	ExchangeKraken        Exchange = "kraken"
	ExchangeKucoin        Exchange = "kucoin"
	ExchangeMEXC          Exchange = "mexc"
	ExchangeOKX           Exchange = "okx"
	ExchangePhemex        Exchange = "phemex"
	ExchangePoloniex      Exchange = "poloniex"
)

// String returns the string representation of the exchange
func (e Exchange) String() string {
	return string(e)
}

// IsValid checks if the exchange is valid
func (e Exchange) IsValid() bool {
	switch e {
	case ExchangeBinance, ExchangeBinanceUS, ExchangeBinanceUSDM,
		ExchangeBitfinex, ExchangeBitget, ExchangeBitmex,
		ExchangeBitstamp, ExchangeBybit, ExchangeCoinbase,
		ExchangeCryptoCom, ExchangeGateIO, ExchangeHuobi,
		ExchangeKraken, ExchangeKucoin, ExchangeMEXC,
		ExchangeOKX, ExchangePhemex, ExchangePoloniex:
		return true
	}
	return false
}
