package taapi

import (
	"fmt"
	"net/http"
)

// Error represents a TAAPI error
type Error struct {
	Message    string
	StatusCode int
	Response   map[string]interface{}
	Err        error
}

// Error implements the error interface
func (e *Error) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("taapi error [%d]: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("taapi error: %s", e.Message)
}

// Unwrap returns the underlying error
func (e *Error) Unwrap() error {
	return e.Err
}

// InvalidArgumentError creates an error for invalid arguments
func InvalidArgumentError(message string) *Error {
	return &Error{
		Message: message,
	}
}

// APIError creates an error from an API response
func APIError(statusCode int, message string, response map[string]interface{}) *Error {
	return &Error{
		Message:    message,
		StatusCode: statusCode,
		Response:   response,
	}
}

// NetworkError creates an error for network issues
func NetworkError(message string, err error) *Error {
	return &Error{
		Message: fmt.Sprintf("network error: %s", message),
		Err:     err,
	}
}

// RateLimitError represents a rate limit error
type RateLimitError struct {
	Message    string
	StatusCode int
	Response   map[string]interface{}
	RetryAfter int
}

// Error implements the error interface
func (e *RateLimitError) Error() string {
	return fmt.Sprintf("taapi rate limit error [%d]: %s (retry after %d seconds)", e.StatusCode, e.Message, e.RetryAfter)
}

// NewRateLimitError creates a new rate limit error
func NewRateLimitError(message string, retryAfter int, response map[string]interface{}) *RateLimitError {
	return &RateLimitError{
		Message:    message,
		StatusCode: http.StatusTooManyRequests,
		Response:   response,
		RetryAfter: retryAfter,
	}
}

// IsRateLimitError checks if an error is a rate limit error
func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	return ok
}

// IsAPIError checks if an error is an API error
func IsAPIError(err error) bool {
	_, ok := err.(*Error)
	return ok
}
