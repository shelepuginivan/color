package gradient

import (
	"image"
	gocolor "image/color"
	"image/draw"

	"github.com/shelepuginivan/color"
	"github.com/shelepuginivan/color/internal/integers"
)

type DiamondGradient struct {
	stops      []*ColorStop
	center     *point
	colorspace Colorspace
}

func NewDiamond(options ...GradientOption) (*DiamondGradient, error) {
	opts := &gradientOptions{}

	for _, opt := range options {
		opt(opts)
	}

	if err := finalizeOptions(opts); err != nil {
		return nil, err
	}

	return &DiamondGradient{
		stops:      opts.stops,
		center:     opts.center,
		colorspace: opts.colorspace,
	}, nil
}

func (dg *DiamondGradient) Colors(steps int) []color.Color {
	return dg.colorspace.Colors(dg.stops, steps)
}

func (dg *DiamondGradient) Render(img image.Image) {
	rect := img.Bounds()
	width := rect.Max.X - rect.Min.Y
	height := rect.Max.Y - rect.Min.Y

	center := dg.center
	if center == nil {
		center = &point{
			x: width / 2,
			y: height / 2,
		}
	}

	maxDist := dg.calcMaxDistance(center, rect)
	colors := dg.Colors(maxDist + 1)

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			dx := integers.Abs(x - center.x)
			dy := integers.Abs(y - center.y)
			dist := dx + dy

			rgb := colors[dist].RGB()
			native := gocolor.RGBA{rgb.R, rgb.G, rgb.B, 255}
			img.(draw.Image).Set(x, y, native)
		}
	}
}

func (dg *DiamondGradient) calcMaxDistance(center *point, rect image.Rectangle) int {
	corners := [][2]int{
		{rect.Min.X, rect.Min.Y},
		{rect.Max.X, rect.Min.Y},
		{rect.Min.X, rect.Max.Y},
		{rect.Max.X, rect.Max.Y},
	}

	res := -1

	for _, c := range corners {
		dx := integers.Abs(c[0] - center.x)
		dy := integers.Abs(c[1] - center.y)

		distance := dx + dy
		res = max(res, distance)
	}

	return res
}
