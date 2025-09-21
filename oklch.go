package color

import (
	"fmt"
	"math"

	"github.com/shelepuginivan/color/internal/degrees"
)

// Oklch represents a color in [Oklch] colorspace, a cylindrical counterpart of
// [Oklab].
//
// [Oklch]: https://en.wikipedia.org/wiki/Oklab_color_space#Conversion_to_and_from_Oklch
type Oklch struct {
	L float64 // L represents lightness of the color.
	C float64 // C represents relative saturation (chroma).
	H int     // H represents angle of the hue in the Oklab color wheel (in degrees).
}

// CMYK returns [CMYK] representation of color (cyan, magenta, yellow, key).
func (c Oklch) CMYK() *CMYK {
	return c.Oklab().CMYK()
}

// Hex returns hexadecimal representation of color.
func (c Oklch) Hex() string {
	return c.Oklab().Hex()
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c Oklch) HSL() *HSL {
	return c.Oklab().HSL()
}

// HSV returns [HSV] representation of color (hue, saturation, value).
func (c Oklch) HSV() *HSV {
	return c.Oklab().HSV()
}

// Lab returns [Lab] representation of color (lightness, red-green,
// yellow-blue).
func (c Oklch) Lab() *Lab {
	return c.Oklab().Lab()
}

// Lch returns [Lch] representation of color (lightness, chroma, hue).
func (c Oklch) Lch() *Lch {
	return c.Oklab().Lch()
}

// Oklab returns [Oklab] representation of color (lightness, red-green,
// yellow-blue).
func (c Oklch) Oklab() *Oklab {
	radH := degrees.ToRadians(c.H)

	return &Oklab{
		L: c.L,
		A: c.C * math.Cos(radH),
		B: c.C * math.Sin(radH),
	}
}

// Oklch returns the color unchanged. This method is required to implement the
// [Color] interface.
func (c Oklch) Oklch() *Oklch {
	return &c
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c Oklch) RGB() *RGB {
	return c.Oklab().RGB()
}

func (c Oklch) XYZ() *XYZ {
	return c.Oklab().XYZ()
}

// Edit allows in-place modification of the [Oklch] color instance using the
// provided editing function.
//
// The returned value is a pointer to the same instance of [Oklch], so it
// should not be used to assign values to other variables. It is intended for
// method chaining.
func (c *Oklch) Edit(editfn func(*Oklch)) *Oklch {
	editfn(c)
	return c
}

// String returns string representation of [Oklab]
func (c Oklch) String() string {
	return fmt.Sprintf("oklch(%.4f, %.4f, %d)", c.L, c.C, c.H)
}
