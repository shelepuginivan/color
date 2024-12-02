package color

// Lab represents a color in [Lab] colorspace.
//
// [Lab]: https://en.wikipedia.org/wiki/CIELAB_color_space
type Lab struct {
	L float64 // L represents lightness of the color.
	A float64 // A represents the green-red component of the color.
	B float64 // B represents the yellow-blue component of the color.
}
