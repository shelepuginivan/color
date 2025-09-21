// Package color implements common color math functions and provides color
// manipulation capabilities.
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
	Oklab() *Oklab
	RGB() *RGB
	XYZ() *XYZ
	String() string
}
