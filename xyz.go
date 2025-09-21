package color

import (
	"fmt"
	"math"
)

// Predefined whitepoints.
var (
	// CIE standard illuminant A. Simulates typical, domestic,
	// tungsten-filament lighting with correlated color temperature of 2856 K.
	A = &XYZ{1.0985, 1.0000, 0.3558}

	// CIE standard illuminant C. Simulates average or north sky daylight with
	// correlated color temperature of 6774 K. Deprecated by CIE.
	C = &XYZ{0.9807, 1.0000, 1.1822}

	// Equal-energy radiator. Useful as a theoretical reference.
	E = &XYZ{1.0000, 1.0000, 1.0000}

	// CIE standard illuminant D50. Simulates warm daylight at sunrise or
	// sunset with correlated color temperature of 5003 K. Also known as
	// horizon light.
	D50 = &XYZ{0.9642, 1.0000, 0.8251}

	// CIE standard illuminant D55. Simulates mid-morning or mid-afternoon
	// daylight with correlated color temperature of 5500 K.
	D55 = &XYZ{0.9568, 1.0000, 0.9214}

	// CIE standard illuminant D65. Simulates noon daylight with correlated
	// color temperature of 6504 K.
	D65 = &XYZ{0.9505, 1.0000, 1.0888}

	// Profile Connection Space (PCS) illuminant used in ICC profiles.
	ICC = &XYZ{0.9642, 1.000, 0.8249}
)

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
	// NOTE: See links below for conversion matrix:
	//   - http://www.brucelindbloom.com/index.html?Eqn_XYZ_to_RGB.html
	//   - https://www.oceanopticsbook.info/view/photometry-and-visibility/from-xyz-to-rgb
	var (
		rVec = 3.2404542*c.X - 1.5371385*c.Y - 0.4985314*c.Z
		gVec = -0.9692660*c.X + 1.8760108*c.Y + 0.0415560*c.Z
		bVec = 0.0556434*c.X - 0.2040259*c.Y + 1.0572252*c.Z
	)

	// Convert from linear form to sRGB.
	var (
		sR = linearToSRGB(rVec)
		sG = linearToSRGB(gVec)
		sB = linearToSRGB(bVec)
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
// [D65] is used as a reference white. Use [XYZ.LabWithWhitepoint] to specify a
// different whitepoint.
func (c XYZ) Lab() *Lab {
	return c.LabWithWhitepoint(D65)
}

// Lch returns [Lch] representation of color (lightness, chroma, hue).
//
// [D65] is used as a reference white. Use [XYZ.LabWithWhitepoint] to specify a
// different whitepoint.
func (c XYZ) Lch() *Lch {
	return c.XYZ().Lab().Lch()
}

// LabWithWhitepoint returns [Lab] representation of color, allowing to
// specify reference white color.
func (c XYZ) LabWithWhitepoint(white *XYZ) *Lab {
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

	return &Lab{l, a, b}
}

// Oklab returns [Oklab] representation of color (lightness, red-green,
// yellow-blue).
func (c XYZ) Oklab() *Oklab {
	// SEE: https://bottosson.github.io/posts/oklab/

	// Convert XYZ to LMS colorspace.
	var (
		l = 0.8189330101*c.X + 0.3618667424*c.Y - 0.1288597137*c.Z
		m = 0.0329845436*c.X + 0.9293118715*c.Y + 0.0361456387*c.Z
		s = 0.0482003018*c.X + 0.2643662691*c.Y + 0.6338517070*c.Z
	)

	// Apply cube root non-linearity.
	var (
		l_ = math.Pow(l, 1.0/3.0)
		m_ = math.Pow(m, 1.0/3.0)
		s_ = math.Pow(s, 1.0/3.0)
	)

	// Convert to Oklab.
	var (
		rL = 0.2104542553*l_ + 0.7936177850*m_ - 0.0040720468*s_
		rA = 1.9779984951*l_ - 2.4285922050*m_ + 0.4505937099*s_
		rB = 0.0259040371*l_ + 0.7827717662*m_ - 0.8086757660*s_
	)

	return &Oklab{
		L: rL,
		A: rA,
		B: rB,
	}
}

// Oklch returns [Oklch] representation of color (lightness, chroma, hue).
func (c XYZ) Oklch() *Oklch {
	return c.XYZ().Oklab().Oklch()
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
