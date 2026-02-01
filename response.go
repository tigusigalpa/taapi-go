package taapi

import "encoding/json"

// IndicatorResponse represents a response for a single indicator
type IndicatorResponse struct {
	Indicator string                 `json:"indicator,omitempty"`
	ID        string                 `json:"id,omitempty"`
	Data      map[string]interface{} `json:"-"`
}

// UnmarshalJSON implements custom JSON unmarshaling
func (r *IndicatorResponse) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if indicator, ok := raw["indicator"].(string); ok {
		r.Indicator = indicator
		delete(raw, "indicator")
	}

	if id, ok := raw["id"].(string); ok {
		r.ID = id
		delete(raw, "id")
	}

	r.Data = raw
	return nil
}

// MarshalJSON implements custom JSON marshaling
func (r *IndicatorResponse) MarshalJSON() ([]byte, error) {
	result := make(map[string]interface{})
	
	if r.Indicator != "" {
		result["indicator"] = r.Indicator
	}
	
	if r.ID != "" {
		result["id"] = r.ID
	}
	
	for k, v := range r.Data {
		result[k] = v
	}
	
	return json.Marshal(result)
}

// GetValue returns the main value from the response
func (r *IndicatorResponse) GetValue() interface{} {
	if val, ok := r.Data["value"]; ok {
		return val
	}
	return r.Data
}

// GetFloat returns a float64 value from the response
func (r *IndicatorResponse) GetFloat(key string) (float64, bool) {
	if val, ok := r.Data[key]; ok {
		if f, ok := val.(float64); ok {
			return f, true
		}
	}
	return 0, false
}

// GetString returns a string value from the response
func (r *IndicatorResponse) GetString(key string) (string, bool) {
	if val, ok := r.Data[key]; ok {
		if s, ok := val.(string); ok {
			return s, true
		}
	}
	return "", false
}

// Get returns a value from the response data
func (r *IndicatorResponse) Get(key string) (interface{}, bool) {
	val, ok := r.Data[key]
	return val, ok
}

// Has checks if a key exists in the response data
func (r *IndicatorResponse) Has(key string) bool {
	_, ok := r.Data[key]
	return ok
}

// BulkResponse represents a response for bulk requests
type BulkResponse struct {
	Responses []*IndicatorResponse
}

// UnmarshalJSON implements custom JSON unmarshaling for bulk responses
func (b *BulkResponse) UnmarshalJSON(data []byte) error {
	var raw []map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	b.Responses = make([]*IndicatorResponse, 0, len(raw))
	for _, item := range raw {
		itemData, err := json.Marshal(item)
		if err != nil {
			continue
		}

		var response IndicatorResponse
		if err := json.Unmarshal(itemData, &response); err != nil {
			continue
		}

		b.Responses = append(b.Responses, &response)
	}

	return nil
}

// FindByID finds a response by its ID
func (b *BulkResponse) FindByID(id string) *IndicatorResponse {
	for _, response := range b.Responses {
		if response.ID == id {
			return response
		}
	}
	return nil
}

// FilterByIndicator returns all responses for a specific indicator
func (b *BulkResponse) FilterByIndicator(indicator string) []*IndicatorResponse {
	var results []*IndicatorResponse
	for _, response := range b.Responses {
		if response.Indicator == indicator {
			results = append(results, response)
		}
	}
	return results
}

// Count returns the number of responses
func (b *BulkResponse) Count() int {
	return len(b.Responses)
}

// Candle represents OHLCV candle data
type Candle struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
}

// ToArray converts a candle to an array format [timestamp, open, high, low, close, volume]
func (c *Candle) ToArray() []interface{} {
	return []interface{}{c.Timestamp, c.Open, c.High, c.Low, c.Close, c.Volume}
}
