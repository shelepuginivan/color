package color

import (
	"fmt"
	"math"

	"github.com/shelepuginivan/color/internal/normalize"
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

// Hex returns hexadecimal representation of color.
func (c RGB) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c RGB) HSL() *HSL {
	var (
		r = float64(c.R) / 255
		g = float64(c.G) / 255
		b = float64(c.B) / 255
	)

	var (
		mx      = max(r, g, b)
		mn      = min(r, g, b)
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
		s = d / (1 - max(2*l-1, 1-2*l))
	}

	return &HSL{
		H: normalize.Degrees(int(math.Round(h))),
		S: normalize.FloatPercents(s),
		L: normalize.FloatPercents(l),
	}
}

// HSV returns [HSV] representation of color (hue, saturation, value).
func (c RGB) HSV() *HSV {
	var (
		r = float64(c.R) / 255
		g = float64(c.G) / 255
		b = float64(c.B) / 255
	)

	var (
		mx      = max(r, g, b)
		mn      = min(r, g, b)
		h, s, v = 0.0, 0.0, mx
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
		s = d / mx
	}

	return &HSV{
		H: normalize.Degrees(int(math.Round(h))),
		S: normalize.FloatPercents(s),
		V: normalize.FloatPercents(v),
	}
}

// XYZ returns [XYZ] representation of color (long wavelengths, brightness,
// short wavelengths).
func (c RGB) XYZ() *XYZ {
	// Convert to sRGB.
	var (
		r = float64(c.R) / 255.0
		g = float64(c.G) / 255.0
		b = float64(c.B) / 255.0
	)

	// Vectorize sRGB values.
	var (
		rVec = srgbToLinear(r) * 100.0
		gVec = srgbToLinear(g) * 100.0
		bVec = srgbToLinear(b) * 100.0
	)

	var (
		x = rVec*0.4124564 + gVec*0.3575761 + bVec*0.1804375
		y = rVec*0.2126729 + gVec*0.7151522 + bVec*0.0721750
		z = rVec*0.0193339 + gVec*0.1191920 + bVec*0.9503041
	)

	return &XYZ{x, y, z}
}

// Lab returns [Lab] representation of color (lightness, red-green,
// yellow-blue).
//
// (95.047, 100.000, 108.883) is used as a reference white. Convert [RGB] to
// [XYZ] with [RGB.XYZ] and use [XYZ.LabWithReferenceWhite] to specify a
// different reference white.
func (c RGB) Lab() *Lab {
	return c.XYZ().Lab()
}

// String returns string representation of [RGB].
func (c RGB) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}

// Edit allows in-place modification of the [RGB] color instance using the
// provided editing function.
//
// The returned value is a pointer to the same instance of [RGB], so it should
// not be used to assign values to other variables. It is intended for method
// chaining.
func (c *RGB) Edit(editfn func(c *RGB)) *RGB {
	editfn(c)
	return c
}
