package taapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntervalString(t *testing.T) {
	assert.Equal(t, "1m", Interval1m.String())
	assert.Equal(t, "1h", Interval1h.String())
	assert.Equal(t, "1d", Interval1d.String())
}

func TestIntervalIsValid(t *testing.T) {
	assert.True(t, Interval1m.IsValid())
	assert.True(t, Interval1h.IsValid())
	assert.True(t, Interval1d.IsValid())
	
	invalidInterval := Interval("invalid")
	assert.False(t, invalidInterval.IsValid())
}

func TestIntervalDuration(t *testing.T) {
	tests := []struct {
		interval Interval
		expected time.Duration
	}{
		{Interval1m, time.Minute},
		{Interval5m, 5 * time.Minute},
		{Interval15m, 15 * time.Minute},
		{Interval30m, 30 * time.Minute},
		{Interval1h, time.Hour},
		{Interval2h, 2 * time.Hour},
		{Interval4h, 4 * time.Hour},
		{Interval12h, 12 * time.Hour},
		{Interval1d, 24 * time.Hour},
		{Interval1w, 7 * 24 * time.Hour},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.interval.Duration())
	}
}
