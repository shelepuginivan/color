package color

// Lab represents a color in [Lab] colorspace.
//
// [Lab]: https://en.wikipedia.org/wiki/CIELAB_color_space
type Lab struct {
	L float64 // L represents lightness of the color.
	A float64 // A represents the green-red component of the color.
	B float64 // B represents the yellow-blue component of the color.
}

// XYZ returns [XYZ] representation of color (long wavelengths, brightness,
// short wavelengths).
//
// (95.047, 100.000, 108.883) is used as a reference white. Use
// [Lab.XYZWithReferenceWhite] to specify a different reference white.
func (c Lab) XYZ() *XYZ {
	// Observer = 2°, Illuminant = D65.
	return c.XYZWithReferenceWhite(&XYZ{
		X: ReferenceWhiteX,
		Y: ReferenceWhiteY,
		Z: ReferenceWhiteZ,
	})
}

// XYZWithReferenceWhite returns [XYZ] representation of color, allowing to
// specify reference white color.
func (c Lab) XYZWithReferenceWhite(white *XYZ) *XYZ {
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

	return &XYZ{
		X: x,
		Y: y,
		Z: z,
	}
}
