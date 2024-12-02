package color

// Default reference white constants. These values determine target white that
// represents "white".
const (
	ReferenceWhiteX = 95.047
	ReferenceWhiteY = 100.000
	ReferenceWhiteZ = 108.883
)

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

	return &Lab{
		L: 116*fy - 16,
		A: 500 * (fx - fy),
		B: 200 * (fy - fz),
	}
}
