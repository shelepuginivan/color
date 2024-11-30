package color

import (
	"math"

	"github.com/shelepuginivan/color/internal/normalize"
)

// [HSL] representation of color.
//
// [HSL]: https://en.wikipedia.org/wiki/HSL
type HSL struct {
	H int // Hue (in degrees).
	S int // Saturation (in percents).
	L int // Lightness (in percents).
}

// NewHSL returns a new instance of [HSL].
func NewHSL(h, s, l int) *HSL {
	return &HSL{
		H: normalize.Degrees(h),
		S: normalize.Percents(s),
		L: normalize.Percents(l),
	}
}

// HSV returns [HSV] representation of color (hue, saturation, value).
func (c HSL) HSV() *HSV {
	// Value normalization.
	var (
		h = normalize.DegreesFloat(c.H)
		s = normalize.PercentsFloat(c.S)
		l = normalize.PercentsFloat(c.L)
	)

	// H_V = H_L
	// S_V = 0 if V equals 0, 2(1 - L/V) otherwise
	// V = L + S_L * min(L, 1-l)
	var (
		hue        = h
		saturation = 0.0
		value      = l + s*min(l, 1-l)
	)

	if value != 0 {
		saturation = 2 * (1 - l/value)
	}

	return &HSV{
		H: normalize.FloatDegrees(hue),
		S: normalize.FloatPercents(saturation),
		V: normalize.FloatPercents(value),
	}
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c HSL) RGB() *RGB {
	var (
		h = normalize.DegreesFloat(c.H)
		s = normalize.PercentsFloat(c.S)
		l = normalize.PercentsFloat(c.L)
	)

	// The default case is when the color is achromatic.
	var (
		r = l
		g = l
		b = l
	)

	if s != 0 {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		r = hueToRGB(p, q, h+1.0/3)
		g = hueToRGB(p, q, h)
		b = hueToRGB(p, q, h-1.0/3)
	}

	return &RGB{
		R: uint8(math.Round(r * 255)),
		G: uint8(math.Round(g * 255)),
		B: uint8(math.Round(b * 255)),
	}
}

// hueToRGB calculates the RGB value for a given hue component. It takes three
// parameters: p and q are the intermediate values calculated from the HSL
// representation, and t represents the normalized hue value (ranging from 0 to
// 1). The function adjusts the value of t to ensure it falls within the range
// [0, 1] by wrapping around if necessary. It then computes the RGB value based
// on the hue's position in the color wheel, returning the corresponding
// float64 value for the red, green, or blue component. This function is used
// internally in the HSL to RGB conversion process.
func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6 {
		return p + (q-p)*6*t
	}
	if t < 0.5 {
		return q
	}
	if t < 2.0/3 {
		return p + (q-p)*(2.0/3-t)*6
	}
	return p
}
