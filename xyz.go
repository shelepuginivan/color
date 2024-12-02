package color

// XYZ represents a color in [XYZ] color space.
//
// [XYZ]: https://en.wikipedia.org/wiki/CIE_1931_color_space
type XYZ struct {
	X float64 // X represents a combination of long wavelengths (red).
	Y float64 // Y corresponds to the luminance or brightness of the color.
	Z float64 // Z captures the short wavelengths (blue).
}

// Lab returns [Lab] representation of color (lightness, red-green,
// yellow-blue).
func (c *XYZ) Lab() *Lab {
	// Observer = 2°, Illuminant = D65.
	var (
		fx = xyzValToLabF(c.X / 95.047)
		fy = xyzValToLabF(c.Y / 100.000)
		fz = xyzValToLabF(c.Z / 108.883)
	)

	return &Lab{
		L: 116*fy - 16,
		A: 500 * (fx - fy),
		B: 200 * (fy - fz),
	}
}
