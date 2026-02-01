package taapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient("test_secret")
	assert.NotNil(t, client)
	assert.Equal(t, "test_secret", client.apiSecret)
	assert.Equal(t, defaultBaseURL, client.baseURL)
	assert.NotNil(t, client.httpClient)
}

func TestClientSetTimeout(t *testing.T) {
	client := NewClient("test_secret")
	result := client.SetTimeout(60)
	assert.Equal(t, client, result)
}

func TestClientSetBaseURL(t *testing.T) {
	client := NewClient("test_secret")
	result := client.SetBaseURL("https://custom.api.com")
	assert.Equal(t, "https://custom.api.com", client.baseURL)
	assert.Equal(t, client, result)
}

func TestClientExchange(t *testing.T) {
	client := NewClient("test_secret")
	builder := client.Exchange(ExchangeBinance)
	assert.NotNil(t, builder)
	assert.Equal(t, "binance", builder.exchange)
}

func TestClientSymbol(t *testing.T) {
	client := NewClient("test_secret")
	builder := client.Symbol("BTC/USDT")
	assert.NotNil(t, builder)
	assert.Equal(t, "BTC/USDT", builder.symbol)
}

func TestClientInterval(t *testing.T) {
	client := NewClient("test_secret")
	builder := client.Interval(Interval1h)
	assert.NotNil(t, builder)
	assert.Equal(t, "1h", builder.interval)
}

func TestClientIndicator(t *testing.T) {
	client := NewClient("test_secret")
	builder := client.Indicator(IndicatorRSI)
	assert.NotNil(t, builder)
	assert.Equal(t, "rsi", builder.indicator)
}

func TestClientDirect(t *testing.T) {
	client := NewClient("test_secret")
	builder := client.Direct()
	assert.NotNil(t, builder)
	assert.NotNil(t, builder.params)
}

func TestClientBulk(t *testing.T) {
	client := NewClient("test_secret")
	builder := client.Bulk()
	assert.NotNil(t, builder)
	assert.NotNil(t, builder.constructs)
}

func TestClientConstruct(t *testing.T) {
	client := NewClient("test_secret")
	builder := client.Construct(ExchangeBinance, "BTC/USDT", Interval1h)
	assert.NotNil(t, builder)
	assert.Equal(t, "binance", builder.exchange)
	assert.Equal(t, "BTC/USDT", builder.symbol)
	assert.Equal(t, "1h", builder.interval)
}

func TestClientManual(t *testing.T) {
	client := NewClient("test_secret")
	builder := client.Manual(IndicatorEMA)
	assert.NotNil(t, builder)
	assert.Equal(t, "ema", builder.indicator)
}
