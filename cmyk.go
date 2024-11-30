package color

import (
	"math"

	"github.com/shelepuginivan/color/internal/normalize"
)

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
	return &CMYK{
		C: normalize.Percents(c),
		M: normalize.Percents(m),
		Y: normalize.Percents(y),
		K: normalize.Percents(k),
	}
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c CMYK) RGB() *RGB {
	var (
		cyan    = normalize.PercentsFloat(c.C)
		magenta = normalize.PercentsFloat(c.M)
		yellow  = normalize.PercentsFloat(c.Y)
		key     = normalize.PercentsFloat(c.K)
	)

	var (
		r = 255 * (1 - cyan) * (1 - key)
		g = 255 * (1 - magenta) * (1 - key)
		b = 255 * (1 - yellow) * (1 - key)
	)

	return &RGB{
		R: uint8(math.Round(r)),
		G: uint8(math.Round(g)),
		B: uint8(math.Round(b)),
	}
}
