package color

import (
	"fmt"
	"math"
)

// RGB represents a color in [RGB] color space.
//
// [RGB]: https://en.wikipedia.org/wiki/RGB_color_model
type RGB struct {
	R uint8 // Red.
	G uint8 // Green.
	B uint8 // Blue.
}

// NewRGB returns a new instance of [RGB].
func NewRGB(r, g, b uint8) *RGB {
	return &RGB{r, g, b}
}

// Hex returns hexadecimal representation of color.
func (c RGB) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// CMYK returns [CMYK] representation of color (cyan, magenta, yellow, key).
func (c RGB) CMYK() *CMYK {
	var (
		r = float64(c.R) / 255
		g = float64(c.G) / 255
		b = float64(c.B) / 255
	)

	var (
		key = 1 - max(r, g, b)
		d   = 1 - key
	)

	// The default case is when key equals 1, i.e. the color is black.
	var (
		cyan    = 0.0
		magenta = 0.0
		yellow  = 0.0
	)

	if d != 0 {
		cyan = (d - r) / d
		magenta = (d - g) / d
		yellow = (d - b) / d
	}

	return &CMYK{
		C: int(math.Round(cyan * 100)),
		M: int(math.Round(magenta * 100)),
		Y: int(math.Round(yellow * 100)),
		K: int(math.Round(key * 100)),
	}
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c RGB) HSL() *HSL {
	var (
		r = float64(c.R) / 255
		g = float64(c.G) / 255
		b = float64(c.B) / 255
	)

	var (
		mx = max(r, g, b)
		mn = min(r, g, b)
	)

	var (
		h, s, l = 0.0, 0.0, (mx + mn) / 2
		d       = mx - mn
	)

	if d != 0 {
		switch mx {
		case r:
			h = (g - b) / d
		case g:
			h = (b-r)/d + 2
		default:
			h = (r-g)/d + 4
		}

		h *= 60
		if h < 0 {
			h += 360
		}
		s = d / (1 - max(2*l-1, 1-2*l))
	}

	return &HSL{
		H: int(h),
		S: int(s * 100),
		L: int(l * 100),
	}
}

func (c RGB) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}
