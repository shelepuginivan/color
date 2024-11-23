package color

import "math"

// Mix calculates the average color from an arbitrary number of colors.
func Mix(colors ...*Color) *Color {
	var (
		total = float64(len(colors))
		r     = 0.0
		g     = 0.0
		b     = 0.0
	)

	for _, color := range colors {
		r += float64(color.R)
		g += float64(color.G)
		b += float64(color.B)
	}

	return &Color{
		R: uint8(math.Round(r / total)),
		G: uint8(math.Round(g / total)),
		B: uint8(math.Round(b / total)),
	}
}
