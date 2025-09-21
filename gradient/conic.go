package gradient

import (
	"image"
	gocolor "image/color"
	"image/draw"
	"math"

	"github.com/shelepuginivan/color"
	"github.com/shelepuginivan/color/internal/degrees"
)

type ConicGradient struct {
	stops      []*ColorStop
	center     *point
	angle      float64
	colorspace Colorspace
}

func NewConic(options ...GradientOption) (*ConicGradient, error) {
	opts := &gradientOptions{}

	for _, opt := range options {
		opt(opts)
	}

	if err := finalizeOptions(opts); err != nil {
		return nil, err
	}

	return &ConicGradient{
		stops:  opts.stops,
		center: opts.center,

		// CSS conic-gradient baseline is vertical (from center point to the top),
		// angle rotates it clockwise.
		angle:      math.Pi/2 - degrees.ToRadians(opts.angle),
		colorspace: opts.colorspace,
	}, nil
}

func (cg *ConicGradient) Colors(steps int) []color.Color {
	return cg.colorspace.Colors(cg.stops, steps)
}

func (cg *ConicGradient) Render(img image.Image) {
	rect := img.Bounds()
	width := rect.Max.X - rect.Min.Y
	height := rect.Max.Y - rect.Min.Y

	center := cg.center
	if center == nil {
		center = &point{
			x: width / 2,
			y: height / 2,
		}
	}

	depth := 360

	colors := cg.Colors(depth)

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			dx := float64(x - center.x)
			dy := float64(y - center.y)
			angle := math.Atan2(dy, dx) + cg.angle
			angle = math.Mod(angle+2*math.Pi, 2*math.Pi)
			fraction := angle / (2 * math.Pi)
			t := int((fraction * float64(depth))) % depth

			rgb := colors[t].RGB()
			native := &gocolor.RGBA{rgb.R, rgb.G, rgb.B, 255}

			img.(draw.Image).Set(x, y, native)
		}
	}
}
