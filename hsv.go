package color

import (
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

func (c HSV) RGB() *RGB {
	var (
		s = normalize.PercentsFloat(c.S)
		v = normalize.PercentsFloat(c.V)
	)

	var (
		chroma = math.Round(s * v)

		// Helper hue value that determines position on color wheel.
		h       = float64(c.H) / 60.0
		x       = chroma * (1 - max(float64(int(h)%2-1), float64(1-int(h)%2)))
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
