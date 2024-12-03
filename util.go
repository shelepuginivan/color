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

// linearToSRGB converts linear representation of color to its channel value in
// normalized form. Defined in [WCAG 2.2]
//
// [WCAG 2.2]: https://www.w3.org/TR/WCAG/#dfn-relative-luminance
func linearToSRGB(c float64) float64 {
	if c <= 0.0031308 {
		return c * 12.92
	}

	return 1.055*math.Pow(c, 1.0/2.4) - 0.055
}

// xyzValToLabF converts [XYZ] value X, Y, and Z to [Lab] f_x, f_y, and f_z
// respectively.
//
// NOTE: http://www.brucelindbloom.com/index.html?Eqn_XYZ_to_Lab.html
func xyzValToLabF(c float64) float64 {
	if c > 0.008856 {
		return math.Pow(c, 1.0/3.0)
	}

	return (903.3*c + 16) / 116
}

// labFToXyzVal converts [Lab] f_x, f_y, and f_z to [XYZ] X, Y, and Z
// respectively.
//
// NOTE: http://www.brucelindbloom.com/index.html?Eqn_Lab_to_XYZ.html
func labFToXyzVal(c float64) float64 {
	c3 := math.Pow(c, 3.0)

	if c3 > 0.008856 {
		return c3
	}

	return (116*c - 16) / 903.3
}
