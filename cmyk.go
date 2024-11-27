package color

import "math"

// CMYK represents a color in [CMYK] color space.
//
// [CMYK]: https://en.wikipedia.org/wiki/CMYK_color_model
type CMYK struct {
	C int // Cyan (in percents).
	M int // Magenta (in percents).
	Y int // Yellow (in percents).
	K int // Black key (in percents).
}

// NewCMYK returns a new instance of [CMYK].
func NewCMYK(c, m, y, k int) *CMYK {
	// TODO: normalize
	return &CMYK{c, m, y, k}
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c CMYK) RGB() *RGB {
	// TODO: normalize
	cyan := float64(c.C) / 100
	magenta := float64(c.M) / 100
	yellow := float64(c.Y) / 100
	key := float64(c.K) / 100

	r := 255 * (1 - cyan) * (1 - key)
	g := 255 * (1 - magenta) * (1 - key)
	b := 255 * (1 - yellow) * (1 - key)

	return &RGB{
		R: uint8(math.Round(r)),
		G: uint8(math.Round(g)),
		B: uint8(math.Round(b)),
	}
}
