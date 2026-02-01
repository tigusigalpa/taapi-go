package taapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDirectBuilderFluent(t *testing.T) {
	client := NewClient("test_secret")
	
	builder := client.Direct().
		Exchange(ExchangeBinance).
		Symbol("BTC/USDT").
		Interval(Interval1h).
		Indicator(IndicatorRSI)
	
	assert.Equal(t, "binance", builder.exchange)
	assert.Equal(t, "BTC/USDT", builder.symbol)
	assert.Equal(t, "1h", builder.interval)
	assert.Equal(t, "rsi", builder.indicator)
}

func TestDirectBuilderWithParams(t *testing.T) {
	client := NewClient("test_secret")
	
	builder := client.Direct().
		WithParams(map[string]interface{}{"period": 14}).
		WithParam("backtrack", 5)
	
	assert.Equal(t, 14, builder.params["period"])
	assert.Equal(t, 5, builder.params["backtrack"])
}

func TestDirectBuilderBacktrack(t *testing.T) {
	client := NewClient("test_secret")
	
	builder := client.Direct().Backtrack(10)
	assert.Equal(t, 10, builder.params["backtrack"])
}

func TestDirectBuilderValidation(t *testing.T) {
	client := NewClient("test_secret")
	
	builder := client.Direct()
	err := builder.validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "exchange is required")
	
	builder.Exchange(ExchangeBinance)
	err = builder.validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "symbol is required")
	
	builder.Symbol("BTC/USDT")
	err = builder.validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "interval is required")
	
	builder.Interval(Interval1h)
	err = builder.validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "indicator is required")
	
	builder.Indicator(IndicatorRSI)
	err = builder.validate()
	assert.NoError(t, err)
}

func TestConstructBuilderAddIndicator(t *testing.T) {
	client := NewClient("test_secret")
	
	construct := client.Construct(ExchangeBinance, "BTC/USDT", Interval1h).
		AddIndicator(IndicatorRSI, map[string]interface{}{"period": 14, "id": "rsi_1"}).
		AddIndicator(IndicatorMACD, map[string]interface{}{"id": "macd_1"})
	
	assert.Equal(t, 2, len(construct.indicators))
}

func TestConstructBuilderToMap(t *testing.T) {
	client := NewClient("test_secret")
	
	construct := client.Construct(ExchangeBinance, "BTC/USDT", Interval1h).
		AddIndicator(IndicatorRSI, map[string]interface{}{"id": "rsi_1"})
	
	result, err := construct.ToMap()
	require.NoError(t, err)
	
	assert.Equal(t, "binance", result["exchange"])
	assert.Equal(t, "BTC/USDT", result["symbol"])
	assert.Equal(t, "1h", result["interval"])
	assert.NotNil(t, result["indicators"])
}

func TestConstructBuilderValidation(t *testing.T) {
	client := NewClient("test_secret")
	
	construct := client.Construct(ExchangeBinance, "BTC/USDT", Interval1h)
	
	_, err := construct.ToMap()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "at least one indicator is required")
}

func TestBulkBuilderAddConstruct(t *testing.T) {
	client := NewClient("test_secret")
	
	bulk := client.Bulk().
		AddConstruct(
			client.Construct(ExchangeBinance, "BTC/USDT", Interval1h).
				AddIndicator(IndicatorRSI, map[string]interface{}{}),
		).
		AddConstruct(
			client.Construct(ExchangeCoinbase, "ETH/USD", Interval4h).
				AddIndicator(IndicatorEMA, map[string]interface{}{"period": 50}),
		)
	
	assert.Equal(t, 2, len(bulk.constructs))
}

func TestManualBuilderWithCandles(t *testing.T) {
	client := NewClient("test_secret")
	
	candles := [][]interface{}{
		{1609459200, 28923.63, 28923.63, 28923.63, 28923.63, 0.0},
		{1609462800, 29083.37, 29188.78, 28963.64, 29103.37, 1107.05626800},
	}
	
	builder := client.Manual(IndicatorEMA).WithCandles(candles)
	
	assert.Equal(t, 2, len(builder.candles))
}

func TestManualBuilderWithCandleStructs(t *testing.T) {
	client := NewClient("test_secret")
	
	candles := []*Candle{
		{Timestamp: 1609459200, Open: 28923.63, High: 28923.63, Low: 28923.63, Close: 28923.63, Volume: 0.0},
		{Timestamp: 1609462800, Open: 29083.37, High: 29188.78, Low: 28963.64, Close: 29103.37, Volume: 1107.05626800},
	}
	
	builder := client.Manual(IndicatorEMA).WithCandleStructs(candles)
	
	assert.Equal(t, 2, len(builder.candles))
}

func TestManualBuilderWithParams(t *testing.T) {
	client := NewClient("test_secret")
	
	builder := client.Manual(IndicatorEMA).
		WithParams(map[string]interface{}{"period": 50}).
		WithParam("backtrack", 5)
	
	assert.Equal(t, 50, builder.params["period"])
	assert.Equal(t, 5, builder.params["backtrack"])
}
