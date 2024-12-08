package color

// Lch represents a color in [Lch] colorspace.
//
// [Lch]: https://en.wikipedia.org/wiki/CIELAB_color_space#Cylindrical_model
type Lch struct {
	L float64 // L represents lightness of the color.
	C float64 // C represents relative saturation (chroma).
	H int     // H represents angle of the hue in the CIELAB color wheel.
}

// NewLch returns a new instance of [Lch].
func NewLch(l, c float64, h int) *Lch {
	return &Lch{l, c, h}
}
