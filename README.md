# TAAPI Go Library

![TAAPI Go](https://github.com/user-attachments/assets/69dae5e6-56bc-4973-9ea2-920f9eaa7d75)

Modern Go library for taapi.io technical analysis API. Features fluent interface, type-safe enums, comprehensive error handling, and full API coverage. Supports direct indicators, bulk requests, and manual calculations with custom candle data. Production-ready with complete test coverage.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue)](https://golang.org)

Modern, idiomatic Go library for the [taapi.io](https://taapi.io/) technical analysis API.

## Features

- ✨ **Modern Go 1.21+** with generics support
- 🔄 **Fluent Interface** for intuitive request building
- 📦 **Full API Coverage** - GET (Direct), POST (Bulk), and POST (Manual) requests
- 🎯 **Type-Safe** - Strongly typed responses and enums
- 🛡️ **Error Handling** - Custom error types for different scenarios
- ✅ **Well Tested** - Comprehensive test coverage
- 📖 **Fully Documented** - Complete GoDoc documentation
- 🚀 **Zero Dependencies** - Only standard library for core functionality

## Installation

```bash
go get github.com/tigusigalpa/taapi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/tigusigalpa/taapi-go"
)

func main() {
    client := taapi.NewClient("YOUR_API_SECRET")
    
    // Get RSI indicator
    rsi, err := client.
        Exchange(taapi.ExchangeBinance).
        Symbol("BTC/USDT").
        Interval(taapi.Interval1h).
        Indicator(taapi.IndicatorRSI).
        Get()
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("RSI: %v\n", rsi.GetValue())
}
```

## Usage

### GET (Direct) Requests

Get a single indicator value for a specific exchange, symbol, and interval.

#### Basic Example

```go
rsi, err := client.
    Exchange(taapi.ExchangeBinance).
    Symbol("BTC/USDT").
    Interval(taapi.Interval1h).
    Indicator(taapi.IndicatorRSI).
    Get()

if err != nil {
    log.Fatal(err)
}

fmt.Printf("RSI Value: %v\n", rsi.GetValue())
```

#### With Additional Parameters

```go
ema, err := client.
    Exchange(taapi.ExchangeBinance).
    Symbol("BTC/USDT").
    Interval(taapi.Interval1h).
    Indicator(taapi.IndicatorEMA).
    WithParams(map[string]interface{}{"period": 50}).
    Get()

fmt.Printf("EMA(50): %v\n", ema.GetValue())
```

#### MACD Example

```go
macd, err := client.
    Exchange(taapi.ExchangeBinance).
    Symbol("ETH/USDT").
    Interval(taapi.Interval4h).
    Indicator(taapi.IndicatorMACD).
    Get()

if err != nil {
    log.Fatal(err)
}

if valueMACD, ok := macd.GetFloat("valueMACD"); ok {
    fmt.Printf("MACD: %v\n", valueMACD)
}
if valueSignal, ok := macd.GetFloat("valueMACDSignal"); ok {
    fmt.Printf("Signal: %v\n", valueSignal)
}
if valueHist, ok := macd.GetFloat("valueMACDHist"); ok {
    fmt.Printf("Histogram: %v\n", valueHist)
}
```

#### With Backtrack

```go
// Get historical data
rsi, err := client.
    Exchange(taapi.ExchangeBinance).
    Symbol("BTC/USDT").
    Interval(taapi.Interval1h).
    Indicator(taapi.IndicatorRSI).
    Backtrack(5).
    Get()

// Get multiple historical values
rsi, err := client.
    Exchange(taapi.ExchangeBinance).
    Symbol("BTC/USDT").
    Interval(taapi.Interval1h).
    Indicator(taapi.IndicatorRSI).
    Backtracks(10).
    Get()
```

### POST (Bulk) Requests

Execute multiple indicator requests in a single API call for better performance.

#### Basic Bulk Request

```go
results, err := client.Bulk().
    AddConstruct(
        client.Construct(taapi.ExchangeBinance, "BTC/USDT", taapi.Interval1h).
            AddIndicator(taapi.IndicatorRSI, map[string]interface{}{"id": "btc_rsi"}).
            AddIndicator(taapi.IndicatorMACD, map[string]interface{}{"id": "btc_macd"}),
    ).
    AddConstruct(
        client.Construct(taapi.ExchangeBinance, "ETH/USDT", taapi.Interval4h).
            AddIndicator(taapi.IndicatorSMA, map[string]interface{}{"period": 200, "id": "eth_sma"}),
    ).
    Execute()

if err != nil {
    log.Fatal(err)
}

// Access results by ID
btcRsi := results.FindByID("btc_rsi")
fmt.Printf("BTC RSI: %v\n", btcRsi.GetValue())

// Iterate through all results
for _, result := range results.Responses {
    fmt.Printf("%s: %v\n", result.Indicator, result.GetValue())
}
```

#### Advanced Bulk Request

```go
results, err := client.Bulk().
    AddConstruct(
        client.Construct(taapi.ExchangeBinance, "BTC/USDT", taapi.Interval1h).
            AddIndicator(taapi.IndicatorRSI, map[string]interface{}{"period": 14, "id": "rsi_14"}).
            AddIndicator(taapi.IndicatorRSI, map[string]interface{}{"period": 21, "id": "rsi_21"}).
            AddIndicator(taapi.IndicatorEMA, map[string]interface{}{"period": 50, "id": "ema_50"}).
            AddIndicator(taapi.IndicatorEMA, map[string]interface{}{"period": 200, "id": "ema_200"}),
    ).
    AddConstruct(
        client.Construct(taapi.ExchangeCoinbase, "ETH/USD", taapi.Interval4h).
            AddIndicator(taapi.IndicatorBBANDS, map[string]interface{}{"id": "eth_bb"}).
            AddIndicator(taapi.IndicatorSTOCH, map[string]interface{}{"id": "eth_stoch"}),
    ).
    Execute()

// Filter by indicator type
rsiResults := results.FilterByIndicator("rsi")
for _, rsi := range rsiResults {
    fmt.Printf("RSI (%s): %v\n", rsi.ID, rsi.GetValue())
}
```

### POST (Manual) Requests

Calculate indicators using your own candle data.

#### Basic Manual Request

```go
candles := [][]interface{}{
    {int64(1609459200), 28923.63, 28923.63, 28923.63, 28923.63, 0.0},
    {int64(1609462800), 29083.37, 29188.78, 28963.64, 29103.37, 1107.05626800},
    {int64(1609466400), 29103.38, 29152.98, 28980.01, 29050.00, 978.58108600},
    // ... more candles
}

ema, err := client.
    Manual(taapi.IndicatorEMA).
    WithCandles(candles).
    WithParams(map[string]interface{}{"period": 50}).
    Execute()

if err != nil {
    log.Fatal(err)
}

fmt.Printf("EMA(50): %v\n", ema.GetValue())
```

#### Using Candle Structs

```go
candles := []*taapi.Candle{
    {Timestamp: 1609459200, Open: 28923.63, High: 28923.63, Low: 28923.63, Close: 28923.63, Volume: 0.0},
    {Timestamp: 1609462800, Open: 29083.37, High: 29188.78, Low: 28963.64, Close: 29103.37, Volume: 1107.05626800},
    // ... more candles
}

rsi, err := client.
    Manual(taapi.IndicatorRSI).
    WithCandleStructs(candles).
    WithParam("period", 14).
    Execute()

fmt.Printf("RSI: %v\n", rsi.GetValue())
```

## Response Handling

### IndicatorResponse

All single indicator requests return an `*IndicatorResponse`:

```go
response, err := client.
    Exchange(taapi.ExchangeBinance).
    Symbol("BTC/USDT").
    Interval(taapi.Interval1h).
    Indicator(taapi.IndicatorRSI).
    Get()

// Get the main value
value := response.GetValue()

// Get specific fields with type safety
if floatVal, ok := response.GetFloat("value"); ok {
    fmt.Printf("Float value: %v\n", floatVal)
}

if strVal, ok := response.GetString("signal"); ok {
    fmt.Printf("String value: %s\n", strVal)
}

// Check if field exists
if response.Has("value") {
    // ...
}

// Access raw data
rawValue, ok := response.Get("value")
```

### BulkResponse

Bulk requests return a `*BulkResponse`:

```go
results, err := client.Bulk().
    AddConstruct(/* ... */).
    Execute()

// Count results
count := results.Count()

// Find by ID
result := results.FindByID("my_indicator")

// Filter by indicator
rsiResults := results.FilterByIndicator("rsi")

// Iterate
for _, response := range results.Responses {
    fmt.Printf("%s: %v\n", response.Indicator, response.GetValue())
}
```

## Error Handling

The library provides custom error types for different scenarios:

```go
result, err := client.
    Exchange(taapi.ExchangeBinance).
    Symbol("BTC/USDT").
    Interval(taapi.Interval1h).
    Indicator(taapi.IndicatorRSI).
    Get()

if err != nil {
    // Check for rate limit error
    if rateLimitErr, ok := err.(*taapi.RateLimitError); ok {
        fmt.Printf("Rate limit exceeded. Retry after: %d seconds\n", rateLimitErr.RetryAfter)
        return
    }
    
    // Check for API error
    if apiErr, ok := err.(*taapi.Error); ok {
        fmt.Printf("API Error [%d]: %s\n", apiErr.StatusCode, apiErr.Message)
        return
    }
    
    // General error
    log.Fatal(err)
}
```

## Available Types

### Exchanges

```go
taapi.ExchangeBinance
taapi.ExchangeBinanceUS
taapi.ExchangeBinanceUSDM
taapi.ExchangeBitfinex
taapi.ExchangeBitget
taapi.ExchangeBitmex
taapi.ExchangeBitstamp
taapi.ExchangeBybit
taapi.ExchangeCoinbase
taapi.ExchangeCryptoCom
taapi.ExchangeGateIO
taapi.ExchangeHuobi
taapi.ExchangeKraken
taapi.ExchangeKucoin
taapi.ExchangeMEXC
taapi.ExchangeOKX
taapi.ExchangePhemex
taapi.ExchangePoloniex
```

### Intervals

```go
taapi.Interval1m   // "1m"
taapi.Interval5m   // "5m"
taapi.Interval15m  // "15m"
taapi.Interval30m  // "30m"
taapi.Interval1h   // "1h"
taapi.Interval2h   // "2h"
taapi.Interval4h   // "4h"
taapi.Interval12h  // "12h"
taapi.Interval1d   // "1d"
taapi.Interval1w   // "1w"
```

### Indicators

```go
taapi.IndicatorRSI
taapi.IndicatorMACD
taapi.IndicatorEMA
taapi.IndicatorSMA
taapi.IndicatorBBANDS
taapi.IndicatorSTOCH
taapi.IndicatorSTOCHRSI
taapi.IndicatorATR
taapi.IndicatorADX
taapi.IndicatorCCI
// ... and many more
```

See the [indicator.go](indicator.go) file for the complete list.

## Advanced Usage

### Custom Timeout

```go
client := taapi.NewClient("YOUR_API_SECRET")
client.SetTimeout(60 * time.Second)
```

### Custom Base URL (for testing)

```go
client := taapi.NewClient("YOUR_API_SECRET")
client.SetBaseURL("https://custom.api.url")
```

## Testing

Run the test suite:

```bash
go test -v ./...
```

Run tests with coverage:

```bash
go test -v -cover ./...
```

## Examples

See the [examples](examples/) directory for complete working examples:

- [Basic Usage](examples/basic/main.go) - Direct and bulk requests
- [Manual Candles](examples/manual/main.go) - Custom candle data

To run examples:

```bash
export TAAPI_SECRET=your_api_secret_here
go run examples/basic/main.go
go run examples/manual/main.go
```

## Requirements

- Go 1.21 or higher
- Valid taapi.io API secret

## Links

- [taapi.io Official Website](https://taapi.io/)
- [taapi.io Documentation](https://taapi.io/documentation/)
- [taapi.io API Reference](https://taapi.io/documentation/integration/)
- [GitHub Repository](https://github.com/tigusigalpa/taapi-go)

## License

This library is open-sourced software licensed under the [MIT license](LICENSE).

## Author

**Igor Sazonov**
- Email: sovletig@gmail.com
- GitHub: [@tigusigalpa](https://github.com/tigusigalpa)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

If you encounter any issues or have questions, please [open an issue](https://github.com/tigusigalpa/taapi-go/issues) on GitHub.
