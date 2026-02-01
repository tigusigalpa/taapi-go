package taapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchangeString(t *testing.T) {
	assert.Equal(t, "binance", ExchangeBinance.String())
	assert.Equal(t, "coinbase", ExchangeCoinbase.String())
	assert.Equal(t, "kraken", ExchangeKraken.String())
}

func TestExchangeIsValid(t *testing.T) {
	assert.True(t, ExchangeBinance.IsValid())
	assert.True(t, ExchangeCoinbase.IsValid())
	assert.True(t, ExchangeKraken.IsValid())
	
	invalidExchange := Exchange("invalid")
	assert.False(t, invalidExchange.IsValid())
}

func TestAllExchanges(t *testing.T) {
	exchanges := []Exchange{
		ExchangeBinance,
		ExchangeBinanceUS,
		ExchangeBinanceUSDM,
		ExchangeBitfinex,
		ExchangeBitget,
		ExchangeBitmex,
		ExchangeBitstamp,
		ExchangeBybit,
		ExchangeCoinbase,
		ExchangeCryptoCom,
		ExchangeGateIO,
		ExchangeHuobi,
		ExchangeKraken,
		ExchangeKucoin,
		ExchangeMEXC,
		ExchangeOKX,
		ExchangePhemex,
		ExchangePoloniex,
	}

	for _, exchange := range exchanges {
		assert.True(t, exchange.IsValid(), "Exchange %s should be valid", exchange)
	}
}
