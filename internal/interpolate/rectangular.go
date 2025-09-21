// Package interpolate provides helper functions for gradient color
// interpolation.
package interpolate

import "math"

// Rectangular performs interpolation in rectangular colorspace.
func Rectangular(v1, v2, t float64) float64 {
	return v1*(1-t) + v2*t
}

// RectangularUint8 performs interpolation in rectangular colorspace on uint8
// values.
func RectangularUint8(v1, v2 uint8, t float64) uint8 {
	return uint8(math.Round(Rectangular(float64(v1), float64(v2), t)))
}
