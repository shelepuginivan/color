package color

import (
	"fmt"
	"math"

	"github.com/shelepuginivan/color/internal/degrees"
)

// Lab represents a color in [Lab] colorspace.
//
// [Lab]: https://en.wikipedia.org/wiki/CIELAB_color_space
type Lab struct {
	L float64 // L represents lightness of the color.
	A float64 // A represents the green-red component of the color.
	B float64 // B represents the yellow-blue component of the color.
}

// NewLab returns a new instance of [Lab] with the default reference white
// [D65].
func NewLab(l, a, b float64) *Lab {
	return &Lab{l, a, b}
}

// CMYK returns [CMYK] representation of color (cyan, magenta, yellow, key).
func (c Lab) CMYK() *CMYK {
	return c.XYZ().CMYK()
}

// Hex returns hexadecimal representation of color.
func (c Lab) Hex() string {
	return c.XYZ().Hex()
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c Lab) HSL() *HSL {
	return c.XYZ().HSL()
}

// HSV returns [HSV] representation of color (hue, saturation, value).
func (c Lab) HSV() *HSV {
	return c.XYZ().HSV()
}

// Lab returns the color unchanged. This method is required to implement the
// [Color] interface.
func (c Lab) Lab() *Lab {
	return &c
}

// Lch returns [Lch] representation of color (lightness, chroma, hue).
func (c Lab) Lch() *Lch {
	return &Lch{
		L: c.L,
		C: math.Sqrt(c.A*c.A + c.B*c.B),
		H: degrees.FromRadians(math.Atan2(c.B, c.A)),
	}
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c Lab) RGB() *RGB {
	return c.XYZ().RGB()
}

// XYZ returns [XYZ] representation of color. [D65] is used as whitepoint, use
// [Lab.XYZWithWhitepoint] to specify a different whitepoint.
func (c Lab) XYZ() *XYZ {
	return c.XYZWithWhitepoint(D65)
}

// XYZWithWhitepoint returns [XYZ] representation of color and allows to
// specify whitepoint.
func (c Lab) XYZWithWhitepoint(white *XYZ) *XYZ {
	var (
		fy = (c.L + 16) / 116
		fx = c.A/500 + fy
		fz = fy - c.B/200
	)

	var (
		x = labFToXyzVal(fx) * white.X
		y = labFToXyzVal(fy) * white.Y
		z = labFToXyzVal(fz) * white.Z
	)

	return &XYZ{x, y, z}
}

// Edit allows in-place modification of the [Lab] color instance using the
// provided editing function.
//
// The returned value is a pointer to the same instance of [Lab], so it should
// not be used to assign values to other variables. It is intended for method
// chaining.
func (c *Lab) Edit(editfn func(c *Lab)) *Lab {
	editfn(c)
	return c
}

// String returns string representation of [Lab].
func (c Lab) String() string {
	return fmt.Sprintf("lab(%.4f, %.4f, %.4f)", c.L, c.A, c.B)
}
