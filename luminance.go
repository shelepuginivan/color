package color

// Luminance calculates relative luminance of the color, normalized to
// 0 for darkest black and 1 for lightest white.
//
// Relative luminance is calculates as per [WCAG 2.2].
//
// [WCAG 2.2]: https://www.w3.org/TR/WCAG/#dfn-relative-luminance
func Luminance(color Color) float64 {
	return color.XYZ().Y
}

// Contrast calculates contrast ratio between two colors.
func Contrast(c1, c2 Color) float64 {
	var (
		l1        = Luminance(c1)
		l2        = Luminance(c2)
		brightest = max(l1, l2)
		darkest   = min(l1, l2)
	)
	return (brightest + 0.05) / (darkest + 0.05)
}

// ContrastBlackWhite returns black or white depending on luminance of the
// given color.
func ContrastBlackWhite(c1 Color) Color {
	if Luminance(c1) > 0.179 {
		return &RGB{0, 0, 0}
	}
	return &RGB{255, 255, 255}
}
