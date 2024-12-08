package color

// Color is a common interface for all package structures. It is used for
// methods that accept colors in any color space.
type Color interface {
	CMYK() *CMYK
	Hex() string
	HSL() *HSL
	HSV() *HSV
	Lab() *Lab
	Lch() *Lch
	RGB() *RGB
	XYZ() *XYZ
	String() string
}
