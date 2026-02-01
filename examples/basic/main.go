package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tigusigalpa/taapi-go"
)

func main() {
	apiSecret := os.Getenv("TAAPI_SECRET")
	if apiSecret == "" {
		log.Fatal("TAAPI_SECRET environment variable is required")
	}

	client := taapi.NewClient(apiSecret)

	fmt.Println("=== TAAPI Go Library - Basic Usage Examples ===\n")

	// Example 1: Simple RSI Request
	fmt.Println("1. Simple RSI Request:")
	rsi, err := client.
		Exchange(taapi.ExchangeBinance).
		Symbol("BTC/USDT").
		Interval(taapi.Interval1h).
		Indicator(taapi.IndicatorRSI).
		Get()

	if err != nil {
		log.Printf("Error getting RSI: %v\n", err)
	} else {
		fmt.Printf("   RSI Value: %v\n\n", rsi.GetValue())
	}

	// Example 2: EMA with Custom Period
	fmt.Println("2. EMA with Custom Period:")
	ema, err := client.
		Exchange(taapi.ExchangeBinance).
		Symbol("BTC/USDT").
		Interval(taapi.Interval1h).
		Indicator(taapi.IndicatorEMA).
		WithParams(map[string]interface{}{"period": 50}).
		Get()

	if err != nil {
		log.Printf("Error getting EMA: %v\n", err)
	} else {
		fmt.Printf("   EMA(50): %v\n\n", ema.GetValue())
	}

	// Example 3: MACD Request
	fmt.Println("3. MACD Request:")
	macd, err := client.
		Exchange(taapi.ExchangeBinance).
		Symbol("ETH/USDT").
		Interval(taapi.Interval4h).
		Indicator(taapi.IndicatorMACD).
		Get()

	if err != nil {
		log.Printf("Error getting MACD: %v\n", err)
	} else {
		if valueMACD, ok := macd.GetFloat("valueMACD"); ok {
			fmt.Printf("   MACD: %v\n", valueMACD)
		}
		if valueSignal, ok := macd.GetFloat("valueMACDSignal"); ok {
			fmt.Printf("   Signal: %v\n", valueSignal)
		}
		if valueHist, ok := macd.GetFloat("valueMACDHist"); ok {
			fmt.Printf("   Histogram: %v\n\n", valueHist)
		}
	}

	// Example 4: Bulk Request
	fmt.Println("4. Bulk Request:")
	results, err := client.Bulk().
		AddConstruct(
			client.Construct(taapi.ExchangeBinance, "BTC/USDT", taapi.Interval1h).
				AddIndicator(taapi.IndicatorRSI, map[string]interface{}{"id": "btc_rsi"}).
				AddIndicator(taapi.IndicatorEMA, map[string]interface{}{"period": 50, "id": "btc_ema"}),
		).
		AddConstruct(
			client.Construct(taapi.ExchangeBinance, "ETH/USDT", taapi.Interval4h).
				AddIndicator(taapi.IndicatorSMA, map[string]interface{}{"period": 200, "id": "eth_sma"}),
		).
		Execute()

	if err != nil {
		log.Printf("Error executing bulk request: %v\n", err)
	} else {
		fmt.Printf("   Total Results: %d\n", results.Count())

		if btcRsi := results.FindByID("btc_rsi"); btcRsi != nil {
			fmt.Printf("   BTC RSI: %v\n", btcRsi.GetValue())
		}

		if btcEma := results.FindByID("btc_ema"); btcEma != nil {
			fmt.Printf("   BTC EMA(50): %v\n", btcEma.GetValue())
		}

		fmt.Println()
	}

	// Example 5: Error Handling
	fmt.Println("5. Error Handling:")
	_, err = client.
		Exchange(taapi.ExchangeBinance).
		Symbol("INVALID/PAIR").
		Interval(taapi.Interval1h).
		Indicator(taapi.IndicatorRSI).
		Get()

	if err != nil {
		if rateLimitErr, ok := err.(*taapi.RateLimitError); ok {
			fmt.Printf("   Rate limit error: retry after %d seconds\n", rateLimitErr.RetryAfter)
		} else if apiErr, ok := err.(*taapi.Error); ok {
			fmt.Printf("   API error [%d]: %s\n", apiErr.StatusCode, apiErr.Message)
		} else {
			fmt.Printf("   Error: %v\n", err)
		}
	}

	fmt.Println("\n=== Examples Complete ===")
}
