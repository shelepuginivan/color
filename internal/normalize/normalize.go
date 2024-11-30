// Package normalize provides internal normalizes for values in different
// formats.
package normalize

import "math"

// Percents normalizes value in percents.
//   - If v > 100, 100 is returned
//   - If v < 0, 0 is returned
//   - Otherwise, v is returned
func Percents(v int) int {
	if v > 100 {
		return 100
	}

	if v < 0 {
		return 0
	}

	return v
}

// PercentsFloat returns normalized value in percents as decimal fraction.
func PercentsFloat(v int) float64 {
	return float64(Percents(v)) / 100
}

// FloatPercents returns value represented as percents.
func FloatPercents(v float64) int {
	if v > 1 {
		return 100
	}

	if v < 0 {
		return 0
	}

	return int(math.Round(v * 100))
}

// Degrees normalizes value in degrees. v mod 360 is returned.
func Degrees(v int) int {
	return v % 360
}

// DegreesFloat returns normalized value in degrees as decimal fraction.
func DegreesFloat(v int) float64 {
	return float64(Degrees(v)) / 360
}

// FloatDegrees returns normalized decimal fraction value in degrees.
func FloatDegrees(v float64) int {
	return Degrees(int(math.Round(360 * v)))
}
