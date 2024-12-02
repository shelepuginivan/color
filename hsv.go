package color

import (
	"fmt"
	"math"

	"github.com/shelepuginivan/color/internal/normalize"
)

// [HSV] representation of color.
//
// [HSV]: https://en.wikipedia.org/wiki/HSL_and_HSV
type HSV struct {
	H int // Hue (in degrees).
	S int // Saturation (in percents).
	V int // Value (in percents).
}

// NewHSV returns a new instance of [HSV].
func NewHSV(h, s, v int) *HSV {
	return &HSV{
		H: normalize.Degrees(h),
		S: normalize.Percents(s),
		V: normalize.Percents(v),
	}
}

// CMYK returns [CMYK] representation of color (cyan, magenta, yellow, key).
func (c HSV) CMYK() *CMYK {
	return c.RGB().CMYK()
}

// Hex returns hexadecimal representation of color.
func (c HSV) Hex() string {
	rgb := c.RGB()
	return fmt.Sprintf("#%02x%02x%02x", rgb.R, rgb.G, rgb.B)
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c HSV) HSL() *HSL {
	// Value normalization.
	var (
		s = normalize.PercentsFloat(c.S)
		v = normalize.PercentsFloat(c.V)
	)

	// H_L = H_V
	// S_L = 0 if L = 0 or L = 1, (V - L) / min(L, 1 - L) otherwise
	// L = V(1 - S_V / 2)
	var (
		hue        = c.H
		saturation = 0.0
		lightness  = v * (1 - s/2)
	)

	if lightness != 0 && lightness != 1 {
		saturation = (v - lightness) / min(lightness, 1-lightness)
	}

	return &HSL{
		H: hue,
		S: normalize.FloatPercents(saturation),
		L: normalize.FloatPercents(lightness),
	}
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c HSV) RGB() *RGB {
	var (
		s = normalize.PercentsFloat(c.S)
		v = normalize.PercentsFloat(c.V)
	)

	// Helper hue value that determines position on color wheel.
	h := float64(c.H)
	if h < 0 {
		h += 360
	} else if h >= 360 {
		h = math.Mod(h, 360)
	}
	h /= 60

	var (
		chroma  = s * v
		rem     = math.Mod(h, 2)
		x       = chroma * (1 - math.Abs(rem-1))
		m       = v - chroma
		r, g, b float64
	)

	switch int(h) {
	case 0:
		r, g, b = chroma, x, 0
	case 1:
		r, g, b = x, chroma, 0
	case 2:
		r, g, b = 0, chroma, x
	case 3:
		r, g, b = 0, x, chroma
	case 4:
		r, g, b = x, 0, chroma
	case 5:
		r, g, b = chroma, 0, x
	}

	return &RGB{
		R: uint8(math.Round((r + m) * 255)),
		G: uint8(math.Round((g + m) * 255)),
		B: uint8(math.Round((b + m) * 255)),
	}
}

// XYZ returns [XYZ] representation of color (long wavelengths, brightness,
// short wavelengths).
func (c HSV) XYZ() *XYZ {
	return c.RGB().XYZ()
}

// String returns string representation of [HSV].
func (c HSV) String() string {
	return fmt.Sprintf("hsv(%d, %d%%, %d%%)", c.H, c.S, c.V)
}

// Edit allows in-place modification of the [HSV] color instance using the
// provided editing function.
//
// The returned value is a pointer to the same instance of [HSV], so it should
// not be used to assign values to other variables. It is intended for method
// chaining.
func (c *HSV) Edit(editfn func(c *HSV)) *HSV {
	editfn(c)
	return c
}
