package taapi

import "fmt"

// DirectBuilder builds direct GET requests
type DirectBuilder struct {
	client    *Client
	exchange  string
	symbol    string
	interval  string
	indicator string
	params    map[string]interface{}
}

// Exchange sets the exchange
func (b *DirectBuilder) Exchange(exchange Exchange) *DirectBuilder {
	b.exchange = exchange.String()
	return b
}

// Symbol sets the trading pair symbol
func (b *DirectBuilder) Symbol(symbol string) *DirectBuilder {
	b.symbol = symbol
	return b
}

// Interval sets the timeframe interval
func (b *DirectBuilder) Interval(interval Interval) *DirectBuilder {
	b.interval = interval.String()
	return b
}

// Indicator sets the technical indicator
func (b *DirectBuilder) Indicator(indicator Indicator) *DirectBuilder {
	b.indicator = indicator.String()
	return b
}

// WithParams adds multiple parameters
func (b *DirectBuilder) WithParams(params map[string]interface{}) *DirectBuilder {
	if b.params == nil {
		b.params = make(map[string]interface{})
	}
	for k, v := range params {
		b.params[k] = v
	}
	return b
}

// WithParam adds a single parameter
func (b *DirectBuilder) WithParam(key string, value interface{}) *DirectBuilder {
	if b.params == nil {
		b.params = make(map[string]interface{})
	}
	b.params[key] = value
	return b
}

// Backtrack sets the backtrack parameter
func (b *DirectBuilder) Backtrack(backtrack int) *DirectBuilder {
	return b.WithParam("backtrack", backtrack)
}

// Backtracks sets the backtracks parameter
func (b *DirectBuilder) Backtracks(backtracks int) *DirectBuilder {
	return b.WithParam("backtracks", backtracks)
}

// Get executes the request
func (b *DirectBuilder) Get() (*IndicatorResponse, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}

	params := make(map[string]interface{})
	params["exchange"] = b.exchange
	params["symbol"] = b.symbol
	params["interval"] = b.interval

	for k, v := range b.params {
		params[k] = v
	}

	return b.client.doGet("/"+b.indicator, params)
}

func (b *DirectBuilder) validate() error {
	if b.exchange == "" {
		return InvalidArgumentError("exchange is required")
	}
	if b.symbol == "" {
		return InvalidArgumentError("symbol is required")
	}
	if b.interval == "" {
		return InvalidArgumentError("interval is required")
	}
	if b.indicator == "" {
		return InvalidArgumentError("indicator is required")
	}
	return nil
}

// ConstructBuilder builds a construct for bulk requests
type ConstructBuilder struct {
	exchange   string
	symbol     string
	interval   string
	indicators []map[string]interface{}
}

// AddIndicator adds an indicator to the construct
func (b *ConstructBuilder) AddIndicator(indicator Indicator, params map[string]interface{}) *ConstructBuilder {
	indicatorData := map[string]interface{}{
		"indicator": indicator.String(),
	}

	for k, v := range params {
		indicatorData[k] = v
	}

	b.indicators = append(b.indicators, indicatorData)
	return b
}

// ToMap converts the construct to a map
func (b *ConstructBuilder) ToMap() (map[string]interface{}, error) {
	if len(b.indicators) == 0 {
		return nil, InvalidArgumentError("at least one indicator is required")
	}

	return map[string]interface{}{
		"exchange":   b.exchange,
		"symbol":     b.symbol,
		"interval":   b.interval,
		"indicators": b.indicators,
	}, nil
}

// BulkBuilder builds bulk POST requests
type BulkBuilder struct {
	client     *Client
	constructs []map[string]interface{}
}

// AddConstruct adds a construct to the bulk request
func (b *BulkBuilder) AddConstruct(construct *ConstructBuilder) *BulkBuilder {
	constructMap, err := construct.ToMap()
	if err != nil {
		return b
	}
	b.constructs = append(b.constructs, constructMap)
	return b
}

// Execute executes the bulk request
func (b *BulkBuilder) Execute() (*BulkResponse, error) {
	if len(b.constructs) == 0 {
		return nil, InvalidArgumentError("at least one construct is required")
	}

	payload := map[string]interface{}{
		"construct": b.constructs,
	}

	result, err := b.client.doPost("/bulk", payload)
	if err != nil {
		return nil, err
	}

	bulkResp, ok := result.(*BulkResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type")
	}

	return bulkResp, nil
}

// ManualBuilder builds manual POST requests with custom candle data
type ManualBuilder struct {
	client    *Client
	indicator string
	candles   [][]interface{}
	params    map[string]interface{}
}

// WithCandles sets the candle data
func (b *ManualBuilder) WithCandles(candles [][]interface{}) *ManualBuilder {
	b.candles = candles
	return b
}

// WithCandleStructs sets the candle data from Candle structs
func (b *ManualBuilder) WithCandleStructs(candles []*Candle) *ManualBuilder {
	b.candles = make([][]interface{}, len(candles))
	for i, candle := range candles {
		b.candles[i] = candle.ToArray()
	}
	return b
}

// WithParams adds multiple parameters
func (b *ManualBuilder) WithParams(params map[string]interface{}) *ManualBuilder {
	for k, v := range params {
		b.params[k] = v
	}
	return b
}

// WithParam adds a single parameter
func (b *ManualBuilder) WithParam(key string, value interface{}) *ManualBuilder {
	b.params[key] = value
	return b
}

// Execute executes the manual request
func (b *ManualBuilder) Execute() (*IndicatorResponse, error) {
	if len(b.candles) == 0 {
		return nil, InvalidArgumentError("candles are required")
	}

	payload := map[string]interface{}{
		"indicator": b.indicator,
		"candles":   b.candles,
	}

	for k, v := range b.params {
		payload[k] = v
	}

	result, err := b.client.doPost("/manual", payload)
	if err != nil {
		return nil, err
	}

	indicatorResp, ok := result.(*IndicatorResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type")
	}

	return indicatorResp, nil
}
