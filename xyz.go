package color

import (
	"fmt"
	"math"
)

// Default reference white constants. These values determine target white that
// represents "white".
const (
	ReferenceWhiteX = 95.047
	ReferenceWhiteY = 100.000
	ReferenceWhiteZ = 108.883
)

// Default reference white.
var DefaultReferenceWhite = &XYZ{
	X: ReferenceWhiteX,
	Y: ReferenceWhiteY,
	Z: ReferenceWhiteZ,
}

// XYZ represents a color in [XYZ] color space.
//
// [XYZ]: https://en.wikipedia.org/wiki/CIE_1931_color_space
type XYZ struct {
	X float64 // X represents a combination of long wavelengths (red).
	Y float64 // Y corresponds to the luminance or brightness of the color.
	Z float64 // Z captures the short wavelengths (blue).
}

// NewXYZ returns a new instance of [XYZ].
func NewXYZ(x, y, z float64) *XYZ {
	return &XYZ{x, y, z}
}

// CMYK returns [CMYK] representation of color (cyan, magenta, yellow, key).
func (c XYZ) CMYK() *CMYK {
	return c.RGB().CMYK()
}

// Hex returns hexadecimal representation of color.
func (c XYZ) Hex() string {
	rgb := c.RGB()
	return fmt.Sprintf("#%02x%02x%02x", rgb.R, rgb.G, rgb.B)
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c XYZ) HSL() *HSL {
	return c.RGB().HSL()
}

// HSV returns [HSV] representation of color (hue, saturation, value).
func (c XYZ) HSV() *HSV {
	return c.RGB().HSV()
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c XYZ) RGB() *RGB {
	// Convert XYZ to sRGB in linear form.
	// NOTE: See links below for convertion matrix:
	//   - http://www.brucelindbloom.com/index.html?Eqn_XYZ_to_RGB.html
	//   - https://www.oceanopticsbook.info/view/photometry-and-visibility/from-xyz-to-rgb
	var (
		rVec = 3.2404542*c.X - 1.5371385*c.Y - 0.4985314*c.Z
		gVec = -0.9692660*c.X + 1.8760108*c.Y + 0.0415560*c.Z
		bVec = 0.0556434*c.X - 0.2040259*c.Y + 1.0572252*c.Z
	)

	// Convert from linear form to sRGB.
	var (
		sR = linearToSRGB(rVec / 100.0)
		sG = linearToSRGB(gVec / 100.0)
		sB = linearToSRGB(bVec / 100.0)
	)

	// Convert to RGB.
	var (
		r = uint8(math.Round(sR * 255))
		g = uint8(math.Round(sG * 255))
		b = uint8(math.Round(sB * 255))
	)

	return &RGB{r, g, b}
}

// Lab returns [Lab] representation of color (lightness, red-green,
// yellow-blue).
//
// (95.047, 100.000, 108.883) is used as a reference white. Use
// [XYZ.LabWithReferenceWhite] to specify a different reference white.
func (c XYZ) Lab() *Lab {
	// Observer = 2°, Illuminant = D65.
	return c.LabWithReferenceWhite(&XYZ{
		X: ReferenceWhiteX,
		Y: ReferenceWhiteY,
		Z: ReferenceWhiteZ,
	})
}

// LabWithReferenceWhite returns [Lab] representation of color, allowing to
// specify reference white color.
func (c XYZ) LabWithReferenceWhite(white *XYZ) *Lab {
	var (
		fx = xyzValToLabF(c.X / white.X)
		fy = xyzValToLabF(c.Y / white.Y)
		fz = xyzValToLabF(c.Z / white.Z)
	)

	var (
		l = 116*fy - 16
		a = 500 * (fx - fy)
		b = 200 * (fy - fz)
	)

	return &Lab{l, a, b, white}
}

// XYZ returns the color unchanged. This method is required to implement the
// [Color] interface.
func (c XYZ) XYZ() *XYZ {
	return &c
}

// Edit allows in-place modification of the [XYZ] color instance using the
// provided editing function.
//
// The returned value is a pointer to the same instance of [XYZ], so it should
// not be used to assign values to other variables. It is intended for method
// chaining.
func (c *XYZ) Edit(editfn func(c *XYZ)) *XYZ {
	editfn(c)
	return c
}

// String returns string representation of [XYZ].
func (c XYZ) String() string {
	return fmt.Sprintf("xyz(%.4f, %.4f, %.4f)", c.X, c.Y, c.Z)
}
