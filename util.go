package color

import "math"

// srgbToLinear converts normalized color channel value to its linear
// representation. Defined in [WCAG 2.2].
//
// [WCAG 2.2]: https://www.w3.org/TR/WCAG/#dfn-relative-luminance
func srgbToLinear(c float64) float64 {
	if c < 0.04045 {
		return c / 12.92
	}
	base := (c + 0.055) / 1.055
	return math.Pow(base, 2.4)
}

// xyzValToLabF converts [XYZ] value X, Y, and Z to [Lab] f_x, f_y, and f_z
// respectively.
func xyzValToLabF(c float64) float64 {
	if c > 0.008856 {
		return math.Pow(c, 1.0/3.0)
	}

	return (903.3*c + 16) / 116
}
