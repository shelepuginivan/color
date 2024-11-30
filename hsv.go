package color

import "github.com/shelepuginivan/color/internal/normalize"

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
