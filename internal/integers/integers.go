// Package integers provides utilities for integer numbers.
package integers

// Abs returns absolute value of an integer.
func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
