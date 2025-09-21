package color

import (
	"fmt"
	"math"
)

// Oklab represents a color in [Oklab] colorspace.
//
// [Oklab]: https://en.wikipedia.org/wiki/Oklab_color_space
type Oklab struct {
	L float64 // L represents lightness of the color.
	A float64 // A represents the green-red component of the color.
	B float64 // B represents the yellow-blue component of the color.
}

func NewOklab(l, a, b float64) *Oklab {
	return &Oklab{l, a, b}
}

// CMYK returns [CMYK] representation of color (cyan, magenta, yellow, key).
func (c Oklab) CMYK() *CMYK {
	return c.XYZ().CMYK()
}

// Hex returns hexadecimal representation of color.
func (c Oklab) Hex() string {
	return c.XYZ().Hex()
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c Oklab) HSL() *HSL {
	return c.XYZ().HSL()
}

// HSV returns [HSV] representation of color (hue, saturation, value).
func (c Oklab) HSV() *HSV {
	return c.XYZ().HSV()
}

// Lab returns the color unchanged. This method is required to implement the
// [Color] interface.
func (c Oklab) Lab() *Lab {
	return c.XYZ().Lab()
}

// Lch returns [Lch] representation of color (lightness, chroma, hue).
func (c Oklab) Lch() *Lch {
	return c.XYZ().Lch()
}

// Oklab returns the color unchanged. This method is required to implement the
// [Color] interface.
func (c Oklab) Oklab() *Oklab {
	return &c
}

// RGB returns [RGB] representation of color (red, green, blue).
func (c Oklab) RGB() *RGB {
	return c.XYZ().RGB()
}

// XYZ returns [XYZ] representation of color (long wavelengths, brightness,
// short wavelengths).
func (c Oklab) XYZ() *XYZ {
	// SEE: https://bottosson.github.io/posts/oklab/
	var (
		l_ = 1.00000000*c.L + 0.39633779*c.A + 0.21580376*c.B
		m_ = 1.00000001*c.L - 0.10556134*c.A - 0.06385417*c.B
		s_ = 1.00000005*c.L - 0.08948418*c.A - 1.29148554*c.B
	)

	var (
		l = math.Pow(l_, 3.0)
		m = math.Pow(m_, 3.0)
		s = math.Pow(s_, 3.0)
	)

	var (
		x = 1.22701385*l - 0.55779998*m + 0.28125615*s
		y = -0.04058018*l + 1.11225687*m - 0.07167668*s
		z = -0.07638128*l - 0.42148198*m + 1.58616322*s
	)

	return &XYZ{x, y, z}
}

// Edit allows in-place modification of the [Oklab] color instance using the
// provided editing function.
//
// The returned value is a pointer to the same instance of [Oklab], so it
// should not be used to assign values to other variables. It is intended for
// method chaining.
func (c *Oklab) Edit(editfn func(*Oklab)) *Oklab {
	editfn(c)
	return c
}

// String returns string representation of [Oklab].
func (c Oklab) String() string {
	return fmt.Sprintf("oklab(%.4f, %.4f, %.4f)", c.L, c.A, c.B)
}
