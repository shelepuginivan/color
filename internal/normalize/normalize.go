package normalize

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

// Degrees normalizes value in degrees. v mod 360 is returned.
func Degrees(v int) int {
	return v % 360
}

// DegreesFloat returns normalized value in degrees as decimal fraction.
func DegreesFloat(v int) float64 {
	return float64(Degrees(v)) / 360
}
