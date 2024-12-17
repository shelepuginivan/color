package degrees

import "math"

// Normalize normalizes value in degrees. v mod 360 is returned.
func Normalize(v int) int {
	norm := v % 360
	if norm < 0 {
		norm += 360
	}
	return norm
}

// ToFloat returns normalized value in degrees as decimal fraction.
func ToFloat(v int) float64 {
	return float64(Normalize(v)) / 360
}

// FromFloat returns normalized decimal fraction value in degrees.
func FromFloat(v float64) int {
	return Normalize(int(math.Round(360 * v)))
}

// ToRadians returns normalized value in degrees as radians.
func ToRadians(v int) float64 {
	return float64(Normalize(v)) / 180 * math.Pi
}

// FromRadians returns normalized value in radians as degrees.
func FromRadians(v float64) int {
	return Normalize(int(math.Round(v * 180 / math.Pi)))
}

// ToTurn returns angle in turns from degrees.
func ToTurn(v int) float64 {
	return float64(v) / 360
}

// FromTurn returns angle in degrees from turns.
func FromTurn(v float64) int {
	return int(math.Round(v * 360))
}
