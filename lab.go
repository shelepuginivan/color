package color

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
