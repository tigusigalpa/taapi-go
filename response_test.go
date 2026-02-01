package taapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndicatorResponseUnmarshal(t *testing.T) {
	jsonData := `{"indicator":"rsi","value":65.5,"id":"test_id"}`
	
	var response IndicatorResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	
	require.NoError(t, err)
	assert.Equal(t, "rsi", response.Indicator)
	assert.Equal(t, "test_id", response.ID)
	assert.Equal(t, 65.5, response.Data["value"])
}

func TestIndicatorResponseGetValue(t *testing.T) {
	response := &IndicatorResponse{
		Indicator: "rsi",
		Data: map[string]interface{}{
			"value": 70.0,
		},
	}
	
	assert.Equal(t, 70.0, response.GetValue())
}

func TestIndicatorResponseGetFloat(t *testing.T) {
	response := &IndicatorResponse{
		Data: map[string]interface{}{
			"value": 65.5,
		},
	}
	
	value, ok := response.GetFloat("value")
	assert.True(t, ok)
	assert.Equal(t, 65.5, value)
	
	_, ok = response.GetFloat("nonexistent")
	assert.False(t, ok)
}

func TestIndicatorResponseGetString(t *testing.T) {
	response := &IndicatorResponse{
		Data: map[string]interface{}{
			"signal": "buy",
		},
	}
	
	value, ok := response.GetString("signal")
	assert.True(t, ok)
	assert.Equal(t, "buy", value)
}

func TestIndicatorResponseHas(t *testing.T) {
	response := &IndicatorResponse{
		Data: map[string]interface{}{
			"value": 65.5,
		},
	}
	
	assert.True(t, response.Has("value"))
	assert.False(t, response.Has("nonexistent"))
}

func TestIndicatorResponseMarshal(t *testing.T) {
	response := &IndicatorResponse{
		Indicator: "rsi",
		ID:        "test_id",
		Data: map[string]interface{}{
			"value": 65.5,
		},
	}
	
	data, err := json.Marshal(response)
	require.NoError(t, err)
	
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	require.NoError(t, err)
	
	assert.Equal(t, "rsi", result["indicator"])
	assert.Equal(t, "test_id", result["id"])
	assert.Equal(t, 65.5, result["value"])
}

func TestBulkResponseUnmarshal(t *testing.T) {
	jsonData := `[
		{"indicator":"rsi","value":65.5,"id":"rsi_1"},
		{"indicator":"macd","valueMACD":1.5,"id":"macd_1"}
	]`
	
	var response BulkResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	
	require.NoError(t, err)
	assert.Equal(t, 2, response.Count())
}

func TestBulkResponseFindByID(t *testing.T) {
	response := &BulkResponse{
		Responses: []*IndicatorResponse{
			{ID: "rsi_1", Indicator: "rsi"},
			{ID: "macd_1", Indicator: "macd"},
		},
	}
	
	found := response.FindByID("rsi_1")
	assert.NotNil(t, found)
	assert.Equal(t, "rsi", found.Indicator)
	
	notFound := response.FindByID("nonexistent")
	assert.Nil(t, notFound)
}

func TestBulkResponseFilterByIndicator(t *testing.T) {
	response := &BulkResponse{
		Responses: []*IndicatorResponse{
			{Indicator: "rsi"},
			{Indicator: "macd"},
			{Indicator: "rsi"},
		},
	}
	
	filtered := response.FilterByIndicator("rsi")
	assert.Equal(t, 2, len(filtered))
}

func TestCandleToArray(t *testing.T) {
	candle := &Candle{
		Timestamp: 1609459200,
		Open:      28923.63,
		High:      28923.63,
		Low:       28923.63,
		Close:     28923.63,
		Volume:    0.0,
	}
	
	array := candle.ToArray()
	assert.Equal(t, 6, len(array))
	assert.Equal(t, int64(1609459200), array[0])
	assert.Equal(t, 28923.63, array[1])
}
