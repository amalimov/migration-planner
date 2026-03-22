package estimation

import (
	"time"
)

// Calculator encapsulates one specific part of the estimation (e.g. "post migration troubleshooting", "storage migration").
type Calculator interface {
	// Name returns the human-readable name of this calculator, used as the key in Engine results.
	Name() string
	// Keys returns the list of Param keys this calculator depends on.
	Keys() []string
	// Calculate runs the estimation using the provided params and returns an Estimation or an error.
	Calculate(params map[string]Param) (Estimation, error)
}

// Param represents an input for a Calculator (can be either user supplied or discovered)
type Param struct {
	Key   string      // Unique identifier (e.g., "network_bandwidth")
	Value interface{} // The actual value (e.g., 1000, "fast", 0.8)
}

// Estimation is the result of a Calculator run.
// Exactly one of {Duration} or {MinDuration, MaxDuration} will be non-nil.
// Use NewPointEstimation or NewRangedEstimation to construct — never build the struct directly.
type Estimation struct {
	Duration    *time.Duration // non-nil for point estimates
	MinDuration *time.Duration // non-nil for ranged estimates
	MaxDuration *time.Duration // non-nil for ranged estimates
	Reason      string
}

// NewPointEstimation constructs an Estimation for a single-value calculator result.
func NewPointEstimation(d time.Duration, reason string) Estimation {
	return Estimation{Duration: &d, Reason: reason}
}

// NewRangedEstimation constructs an Estimation for a calculator that returns a duration range.
// min must be <= max.
func NewRangedEstimation(min, max time.Duration, reason string) Estimation {
	return Estimation{MinDuration: &min, MaxDuration: &max, Reason: reason}
}

// IsRanged reports whether this estimation carries a range rather than a point value.
func (e Estimation) IsRanged() bool { return e.MinDuration != nil }
