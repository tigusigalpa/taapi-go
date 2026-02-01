package taapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultBaseURL = "https://api.taapi.io"
	defaultTimeout = 30 * time.Second
)

// Client represents a TAAPI API client
type Client struct {
	apiSecret  string
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new TAAPI client
func NewClient(apiSecret string) *Client {
	return &Client{
		apiSecret: apiSecret,
		baseURL:   defaultBaseURL,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// SetTimeout sets the HTTP client timeout
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.httpClient.Timeout = timeout
	return c
}

// SetBaseURL sets a custom base URL (useful for testing)
func (c *Client) SetBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

// Exchange starts building a direct request with an exchange
func (c *Client) Exchange(exchange Exchange) *DirectBuilder {
	return &DirectBuilder{
		client:   c,
		exchange: exchange.String(),
	}
}

// Symbol starts building a direct request with a symbol
func (c *Client) Symbol(symbol string) *DirectBuilder {
	return &DirectBuilder{
		client: c,
		symbol: symbol,
	}
}

// Interval starts building a direct request with an interval
func (c *Client) Interval(interval Interval) *DirectBuilder {
	return &DirectBuilder{
		client:   c,
		interval: interval.String(),
	}
}

// Indicator starts building a direct request with an indicator
func (c *Client) Indicator(indicator Indicator) *DirectBuilder {
	return &DirectBuilder{
		client:    c,
		indicator: indicator.String(),
	}
}

// Direct creates a new direct request builder
func (c *Client) Direct() *DirectBuilder {
	return &DirectBuilder{
		client: c,
		params: make(map[string]interface{}),
	}
}

// Bulk creates a new bulk request builder
func (c *Client) Bulk() *BulkBuilder {
	return &BulkBuilder{
		client:     c,
		constructs: make([]map[string]interface{}, 0),
	}
}

// Construct creates a new construct for bulk requests
func (c *Client) Construct(exchange Exchange, symbol string, interval Interval) *ConstructBuilder {
	return &ConstructBuilder{
		exchange:   exchange.String(),
		symbol:     symbol,
		interval:   interval.String(),
		indicators: make([]map[string]interface{}, 0),
	}
}

// Manual creates a new manual request builder
func (c *Client) Manual(indicator Indicator) *ManualBuilder {
	return &ManualBuilder{
		client:    c,
		indicator: indicator.String(),
		params:    make(map[string]interface{}),
	}
}

// doGet performs a GET request
func (c *Client) doGet(endpoint string, params map[string]interface{}) (*IndicatorResponse, error) {
	urlStr := c.baseURL + endpoint

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, NetworkError("invalid URL", err)
	}

	q := u.Query()
	q.Set("secret", c.apiSecret)

	for key, value := range params {
		q.Set(key, fmt.Sprintf("%v", value))
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NetworkError("failed to create request", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, NetworkError("request failed", err)
	}
	defer resp.Body.Close()

	result, err := c.handleResponse(resp, false)
	if err != nil {
		return nil, err
	}

	indicatorResp, ok := result.(*IndicatorResponse)
	if !ok {
		return nil, APIError(0, "unexpected response type", nil)
	}

	return indicatorResp, nil
}

// doPost performs a POST request
func (c *Client) doPost(endpoint string, payload map[string]interface{}) (interface{}, error) {
	urlStr := c.baseURL + endpoint

	payload["secret"] = c.apiSecret

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, NetworkError("failed to marshal JSON", err)
	}

	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, NetworkError("failed to create request", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, NetworkError("request failed", err)
	}
	defer resp.Body.Close()

	isBulk := endpoint == "/bulk"
	return c.handleResponse(resp, isBulk)
}

// handleResponse processes the HTTP response
func (c *Client) handleResponse(resp *http.Response, isBulk bool) (interface{}, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NetworkError("failed to read response body", err)
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		retryAfter := 0
		if retryAfterStr := resp.Header.Get("Retry-After"); retryAfterStr != "" {
			retryAfter, _ = strconv.Atoi(retryAfterStr)
		}

		var errorData map[string]interface{}
		json.Unmarshal(body, &errorData)

		message := "rate limit exceeded"
		if msg, ok := errorData["error"].(string); ok {
			message = msg
		}

		return nil, NewRateLimitError(message, retryAfter, errorData)
	}

	if resp.StatusCode >= 400 {
		var errorData map[string]interface{}
		json.Unmarshal(body, &errorData)

		message := "unknown API error"
		if msg, ok := errorData["error"].(string); ok {
			message = msg
		} else if msg, ok := errorData["message"].(string); ok {
			message = msg
		}

		return nil, APIError(resp.StatusCode, message, errorData)
	}

	if isBulk {
		var bulkResp BulkResponse
		if err := json.Unmarshal(body, &bulkResp); err != nil {
			return nil, APIError(0, "failed to decode bulk response", nil)
		}
		return &bulkResp, nil
	}

	var indicatorResp IndicatorResponse
	if err := json.Unmarshal(body, &indicatorResp); err != nil {
		return nil, APIError(0, "failed to decode response", nil)
	}
	return &indicatorResp, nil
}
