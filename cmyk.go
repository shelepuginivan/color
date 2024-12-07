package color

import (
	"fmt"
	"math"

	"github.com/shelepuginivan/color/internal/percents"
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
		C: percents.Normalize(c),
		M: percents.Normalize(m),
		Y: percents.Normalize(y),
		K: percents.Normalize(k),
	}
}

// CMYK returns the color unchanged. This method is required to implement the
// [Color] interface.
func (c CMYK) CMYK() *CMYK {
	return &c
}

// Hex returns hexadecimal representation of color.
func (c CMYK) Hex() string {
	rgb := c.RGB()
	return fmt.Sprintf("#%02x%02x%02x", rgb.R, rgb.G, rgb.B)
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c CMYK) HSL() *HSL {
	return c.RGB().HSL()
}

// HSV returns [HSV] representation of color (hue, saturation, value).
func (c CMYK) HSV() *HSV {
	return c.RGB().HSV()
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c CMYK) RGB() *RGB {
	var (
		cyan    = percents.ToFloat(c.C)
		magenta = percents.ToFloat(c.M)
		yellow  = percents.ToFloat(c.Y)
		key     = percents.ToFloat(c.K)
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

// XYZ returns [XYZ] representation of color (long wavelengths, brightness,
// short wavelengths).
func (c CMYK) XYZ() *XYZ {
	return c.RGB().XYZ()
}

// Lab returns [Lab] representation of color (lightness, red-green,
// yellow-blue).
//
// [D65] is used as a reference white. Use [XYZ.LabWithWhitepoint] to specify a
// different whitepoint.
func (c CMYK) Lab() *Lab {
	return c.XYZ().Lab()
}

// Edit allows in-place modification of the [CMYK] color instance using the
// provided editing function.
//
// The returned value is a pointer to the same instance of [CMYK], so it should
// not be used to assign values to other variables. It is intended for method
// chaining.
func (c *CMYK) Edit(editfn func(c *CMYK)) *CMYK {
	editfn(c)
	return c
}

// String returns string representation of [CMYK].
func (c CMYK) String() string {
	return fmt.Sprintf("cmyk(%d%%, %d%%, %d%%, %d%%)", c.C, c.M, c.Y, c.K)
}
