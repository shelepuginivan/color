package gradient

import (
	"image"
	gocolor "image/color"
	"image/draw"

	"github.com/shelepuginivan/color"
)

type RadialGradient struct {
	stops      []*ColorStop
	center     pointSpec
	colorspace Colorspace
}

func NewRadial(options ...GradientOption) (*RadialGradient, error) {
	opts := &gradientOptions{}

	for _, opt := range options {
		opt(opts)
	}

	if err := finalizeOptions(opts); err != nil {
		return nil, err
	}

	return &RadialGradient{
		stops:      opts.stops,
		center:     opts.center,
		colorspace: opts.colorspace,
	}, nil
}

func (rg *RadialGradient) Colors(steps int) []color.Color {
	return rg.colorspace.Colors(rg.stops, steps)
}

func (rg *RadialGradient) Render(img image.Image) {
	rect := img.Bounds()
	center := rg.center.Position(rect)

	maxDist := rg.calcMaxDistance(center, rect)
	colors := rg.Colors(maxDist + 1)

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			dx := x - center.x
			dy := y - center.y
			dist := dx*dx + dy*dy

			rgb := colors[dist].RGB()
			native := gocolor.RGBA{rgb.R, rgb.G, rgb.B, 255}
			img.(draw.Image).Set(x, y, native)
		}
	}
}

func (rg *RadialGradient) calcMaxDistance(center point, rect image.Rectangle) int {
	corners := [][2]int{
		{rect.Min.X, rect.Min.Y},
		{rect.Max.X, rect.Min.Y},
		{rect.Min.X, rect.Max.Y},
		{rect.Max.X, rect.Max.Y},
	}

	res := -1

	for _, c := range corners {
		dx := c[0] - center.x
		dy := c[1] - center.y

		distanceSqr := dx*dx + dy*dy
		res = max(res, distanceSqr)
	}

	return res
}
