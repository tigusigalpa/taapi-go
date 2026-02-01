package taapi

import "time"

// Interval represents a timeframe interval
type Interval string

const (
	Interval1m  Interval = "1m"
	Interval5m  Interval = "5m"
	Interval15m Interval = "15m"
	Interval30m Interval = "30m"
	Interval1h  Interval = "1h"
	Interval2h  Interval = "2h"
	Interval4h  Interval = "4h"
	Interval12h Interval = "12h"
	Interval1d  Interval = "1d"
	Interval1w  Interval = "1w"
)

// String returns the string representation of the interval
func (i Interval) String() string {
	return string(i)
}

// IsValid checks if the interval is valid
func (i Interval) IsValid() bool {
	switch i {
	case Interval1m, Interval5m, Interval15m, Interval30m,
		Interval1h, Interval2h, Interval4h, Interval12h,
		Interval1d, Interval1w:
		return true
	}
	return false
}

// Duration returns the time.Duration representation of the interval
func (i Interval) Duration() time.Duration {
	switch i {
	case Interval1m:
		return time.Minute
	case Interval5m:
		return 5 * time.Minute
	case Interval15m:
		return 15 * time.Minute
	case Interval30m:
		return 30 * time.Minute
	case Interval1h:
		return time.Hour
	case Interval2h:
		return 2 * time.Hour
	case Interval4h:
		return 4 * time.Hour
	case Interval12h:
		return 12 * time.Hour
	case Interval1d:
		return 24 * time.Hour
	case Interval1w:
		return 7 * 24 * time.Hour
	default:
		return 0
	}
}
