// Package percents provides internal normalization functions for percents.
package percents

import "math"

// Normalize normalizes value in percents.
//   - If v > 100, 100 is returned
//   - If v < 0, 0 is returned
//   - Otherwise, v is returned
func Normalize(v int) int {
	if v > 100 {
		return 100
	}

	if v < 0 {
		return 0
	}

	return v
}

// ToFloat returns normalized value in percents as decimal fraction.
func ToFloat(v int) float64 {
	return float64(Normalize(v)) / 100
}

// FromFloat returns value represented as percents.
func FromFloat(v float64) int {
	if v > 1 {
		return 100
	}

	if v < 0 {
		return 0
	}

	return int(math.Round(v * 100))
}
