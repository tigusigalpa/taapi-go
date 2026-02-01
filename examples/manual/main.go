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

	fmt.Println("=== Manual Candles Example ===\n")

	// Sample candle data: [timestamp, open, high, low, close, volume]
	candles := [][]interface{}{
		{int64(1609459200), 28923.63, 28923.63, 28923.63, 28923.63, 0.0},
		{int64(1609462800), 29083.37, 29188.78, 28963.64, 29103.37, 1107.05626800},
		{int64(1609466400), 29103.38, 29152.98, 28980.01, 29050.00, 978.58108600},
		{int64(1609470000), 29050.01, 29071.99, 28852.01, 28852.02, 1223.86114900},
		{int64(1609473600), 28852.02, 29034.99, 28827.00, 28961.00, 1239.48652100},
		{int64(1609477200), 28961.01, 29150.00, 28943.00, 29099.99, 916.25317300},
		{int64(1609480800), 29099.98, 29188.00, 29027.14, 29159.99, 848.81565000},
		{int64(1609484400), 29160.00, 29289.99, 29115.00, 29250.00, 1042.33776900},
		{int64(1609488000), 29250.01, 29377.77, 29217.78, 29377.76, 1277.13690800},
		{int64(1609491600), 29377.77, 29480.00, 29341.00, 29374.99, 1275.91266700},
		{int64(1609495200), 29375.00, 29432.99, 29320.00, 29432.98, 770.12843100},
		{int64(1609498800), 29432.99, 29600.00, 29415.00, 29600.00, 1361.93171000},
		{int64(1609502400), 29600.01, 29617.78, 29470.00, 29546.63, 1095.14694200},
		{int64(1609506000), 29546.64, 29546.64, 29200.00, 29288.00, 1868.26219600},
		{int64(1609509600), 29288.01, 29315.00, 29101.00, 29196.00, 1448.38314900},
	}

	fmt.Printf("Candle data format: [timestamp, open, high, low, close, volume]\n")
	fmt.Printf("Total candles: %d\n\n", len(candles))

	// Example 1: Calculate EMA(10)
	fmt.Println("1. Calculate EMA(10):")
	ema, err := client.
		Manual(taapi.IndicatorEMA).
		WithCandles(candles).
		WithParams(map[string]interface{}{"period": 10}).
		Execute()

	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("   EMA(10): %v\n\n", ema.GetValue())
	}

	// Example 2: Calculate RSI(14)
	fmt.Println("2. Calculate RSI(14):")
	rsi, err := client.
		Manual(taapi.IndicatorRSI).
		WithCandles(candles).
		WithParam("period", 14).
		Execute()

	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("   RSI(14): %v\n\n", rsi.GetValue())
	}

	// Example 3: Using Candle structs
	fmt.Println("3. Using Candle structs:")
	candleStructs := []*taapi.Candle{
		{Timestamp: 1609459200, Open: 28923.63, High: 28923.63, Low: 28923.63, Close: 28923.63, Volume: 0.0},
		{Timestamp: 1609462800, Open: 29083.37, High: 29188.78, Low: 28963.64, Close: 29103.37, Volume: 1107.05626800},
		{Timestamp: 1609466400, Open: 29103.38, High: 29152.98, Low: 28980.01, Close: 29050.00, Volume: 978.58108600},
		{Timestamp: 1609470000, Open: 29050.01, High: 29071.99, Low: 28852.01, Close: 28852.02, Volume: 1223.86114900},
		{Timestamp: 1609473600, Open: 28852.02, High: 29034.99, Low: 28827.00, Close: 28961.00, Volume: 1239.48652100},
		{Timestamp: 1609477200, Open: 28961.01, High: 29150.00, Low: 28943.00, Close: 29099.99, Volume: 916.25317300},
		{Timestamp: 1609480800, Open: 29099.98, High: 29188.00, Low: 29027.14, Close: 29159.99, Volume: 848.81565000},
		{Timestamp: 1609484400, Open: 29160.00, High: 29289.99, Low: 29115.00, Close: 29250.00, Volume: 1042.33776900},
		{Timestamp: 1609488000, Open: 29250.01, High: 29377.77, Low: 29217.78, Close: 29377.76, Volume: 1277.13690800},
		{Timestamp: 1609491600, Open: 29377.77, High: 29480.00, Low: 29341.00, Close: 29374.99, Volume: 1275.91266700},
	}

	sma, err := client.
		Manual(taapi.IndicatorSMA).
		WithCandleStructs(candleStructs).
		WithParam("period", 10).
		Execute()

	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("   SMA(10): %v\n\n", sma.GetValue())
	}

	fmt.Println("=== Example Complete ===")
}
