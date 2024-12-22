package color

// Complementary returns a complementary color for the given one.
// Complementary colors are two colors that are directly across from one
// another on the color wheel.
func Complementary(c Color) Color {
	hsl := c.HSL()
	return NewHSL(hsl.H+180, hsl.S, hsl.L)
}

// SplitComplementary returns two colors adjacent to the complementary color of
// the given one.
func SplitComplementary(c Color) (Color, Color) {
	hsl := c.HSL()
	return NewHSL(hsl.H+150, hsl.S, hsl.L), NewHSL(hsl.H+210, hsl.S, hsl.L)
}

// Triadic returns 2 of the remaining colors from the triadic colorscheme for
// the given color.
func Triadic(c Color) (Color, Color) {
	hsl := c.HSL()
	return NewHSL(hsl.H+120, hsl.S, hsl.L), NewHSL(hsl.H+240, hsl.S, hsl.L)
}

// Tetradic returns 3 of the remaining colors from the tetradic colorscheme for
// the given color.
func Tetradic(c Color) (Color, Color, Color) {
	hsl := c.HSL()
	return NewHSL(hsl.H+90, hsl.S, hsl.L), NewHSL(hsl.H+180, hsl.S, hsl.L), NewHSL(hsl.H+270, hsl.S, hsl.L)
}

// Analogous returns two colors adjacent to the given one.
func Analogous(c Color) (Color, Color) {
	hsl := c.HSL()
	return NewHSL(hsl.H-30, hsl.S, hsl.L), NewHSL(hsl.H+30, hsl.S, hsl.L)
}
