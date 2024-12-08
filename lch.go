package color

import (
	"fmt"
	"math"

	"github.com/shelepuginivan/color/internal/degrees"
)

// Lch represents a color in [Lch] colorspace.
//
// [Lch]: https://en.wikipedia.org/wiki/CIELAB_color_space#Cylindrical_model
type Lch struct {
	L float64 // L represents lightness of the color.
	C float64 // C represents relative saturation (chroma).
	H int     // H represents angle of the hue in the CIELAB color wheel.
}

// NewLch returns a new instance of [Lch].
func NewLch(l, c float64, h int) *Lch {
	return &Lch{l, c, h}
}

// CMYK returns [CMYK] representation of color (cyan, magenta, yellow, key).
func (c Lch) CMYK() *CMYK {
	return c.Lab().CMYK()
}

// Hex returns hexadecimal representation of color.
func (c Lch) Hex() string {
	return c.Lab().Hex()
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c Lch) HSL() *HSL {
	return c.Lab().HSL()
}

// HSV returns [HSV] representation of color (hue, saturation, value).
func (c Lch) HSV() *HSV {
	return c.Lab().HSV()
}

// Lab returns [Lab] representation of color (lightness, red-green,
// yellow-blue).
func (c Lch) Lab() *Lab {
	radH := degrees.ToRadians(c.H)

	return &Lab{
		L: c.L,
		A: c.C * math.Cos(radH),
		B: c.C * math.Sin(radH),
	}
}

// Lch returns the color unchanged. This method is required to implement the
// [Color] interface.
func (c Lch) Lch() *Lch {
	return &c
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c Lch) RGB() *RGB {
	return c.Lab().RGB()
}

// XYZ returns [XYZ] representation of color. [D65] is used as whitepoint, use
// [Lab.XYZWithWhitepoint] to specify a different whitepoint.
func (c Lch) XYZ() *XYZ {
	return c.Lab().XYZ()
}

// Edit allows in-place modification of the [Lch] color instance using the
// provided editing function.
//
// The returned value is a pointer to the same instance of [Lch], so it should
// not be used to assign values to other variables. It is intended for method
// chaining.
func (c *Lch) Edit(editfn func(c *Lch)) *Lch {
	editfn(c)
	return c
}

// String returns string representation of [Lch].
func (c Lch) String() string {
	return fmt.Sprintf("lch(%.4f, %.4f, %d)", c.L, c.C, c.H)
}
