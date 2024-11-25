package color

import "math"

// Luminance coefficients.
const (
	LuminanceRed   = 0.2126
	LuminanceGreen = 0.7152
	LuminanceBlue  = 0.0722
)

// Luminance calculates relative luminance of the color, normalized to
// 0 for darkest black and 1 for lightest white.
//
// Relative luminance is calculates as per [WCAG 2.2].
//
// [WCAG 2.2]: https://www.w3.org/TR/WCAG/#dfn-relative-luminance
func Luminance(c *Color) float64 {
	r := float64(c.R) / 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255

	rLinear := srgbToLinear(r)
	gLinear := srgbToLinear(g)
	bLinear := srgbToLinear(b)

	return LuminanceRed*rLinear + LuminanceGreen*gLinear + LuminanceBlue*bLinear
}

// Contrast calculates contrast ratio between two colors.
func Contrast(c1, c2 *Color) float64 {
	l1 := Luminance(c1)
	l2 := Luminance(c2)

	brightest := max(l1, l2)
	darkest := min(l1, l2)

	return (brightest + 0.05) / (darkest + 0.05)
}

// ContrastBlackWhite returns black or white depending on luminance of the
// given color. An example use case is to determine the text color that
// contrasts better with a given color.
func ContrastBlackWhite(c1 *Color) *Color {
	if Luminance(c1) > 0.179 {
		return &Color{0, 0, 0}
	}
	return &Color{255, 255, 255}
}

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
