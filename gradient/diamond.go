package gradient

import (
	"image"
	gocolor "image/color"
	"image/draw"

	"github.com/shelepuginivan/color"
	"github.com/shelepuginivan/color/internal/integers"
)

// DiamondGradient is a gradient with a progressive color transition radiating
// from the center point in diamond shapes.
type DiamondGradient struct {
	stops      []*ColorStop
	center     pointSpec
	colorspace Colorspace
}

// NewDiamond returns a new instance of [DiamondGradient].
//
// Supported options are:
//   - [WithCenterAt] and [WithRelativeCenter] set the center point of the
//     gradient.
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
	center := dg.center.Position(rect)

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

func (dg *DiamondGradient) calcMaxDistance(center point, rect image.Rectangle) int {
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
