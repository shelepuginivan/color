package color

import "math"

// [HSL] representation of color.
//
// [HSL]: https://en.wikipedia.org/wiki/HSL
type HSL struct {
	Hue        int // Hue (in degrees).
	Saturation int // Saturation (in percents).
	Lightness  int // Lightness (in percents).
}

// NewHSL returns a new instance of [HSL].
func NewHSL(h, s, l int) *HSL {
	// TODO: normalize percents and degrees.
	return &HSL{h, s, l}
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c HSL) RGB() *RGB {
	// TODO: normalize percents and degrees.
	h := float64(c.Hue) / 360
	s := float64(c.Saturation) / 100
	l := float64(c.Lightness) / 100

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
