package color

import "fmt"

// Lab represents a color in [Lab] colorspace.
//
// [Lab]: https://en.wikipedia.org/wiki/CIELAB_color_space
type Lab struct {
	L float64 // L represents lightness of the color.
	A float64 // A represents the green-red component of the color.
	B float64 // B represents the yellow-blue component of the color.

	White *XYZ // Reference white.
}

// NewLab returns a new instance of [Lab] with the default reference white.
func NewLab(l, a, b float64) *Lab {
	return &Lab{l, a, b, DefaultReferenceWhite}
}

// NewLabWithReferenceWhite returns a new instance of [Lab] and allows to set a
// custom reference white.
func NewLabWithReferenceWhite(l, a, b float64, white *XYZ) *Lab {
	return &Lab{l, a, b, white}
}

// SetReferenceWhite sets [Lab] reference white.
func (c *Lab) SetReferenceWhite(white *XYZ) *Lab {
	color := c.XYZ().LabWithReferenceWhite(white)

	c.L = color.L
	c.A = color.A
	c.B = color.B

	return color
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

// RGB returns [RGB] representation of color (red, green, blue).
func (c Lab) RGB() *RGB {
	return c.XYZ().RGB()
}

// XYZWithReferenceWhite returns [XYZ] representation of color.
//
// [Lab.White] is used as a reference white.
func (c Lab) XYZ() *XYZ {
	var (
		fy = (c.L + 16) / 116
		fx = c.A/500 + fy
		fz = fy - c.B/200
	)

	var (
		x = labFToXyzVal(fx) * c.White.X
		y = labFToXyzVal(fy) * c.White.Y
		z = labFToXyzVal(fz) * c.White.Z
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
