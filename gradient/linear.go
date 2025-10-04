package gradient

import (
	"image"
	gocolor "image/color"
	"image/draw"
	"math"

	"github.com/shelepuginivan/color"
)

// LinearGradient is a gradient with a progressive color transition along a
// straight line.
type LinearGradient struct {
	stops      []*ColorStop
	colorspace Colorspace
	angle      angleSpec
}

// NewLinear returns a new instance of [LinearGradient].
//
// Supported options are:
//   - [WithAngle] and [WithDirection] set the angle of gradient line.
//     The default angle is 0° (gradient progresses from bottom to top).
func NewLinear(options ...GradientOption) (*LinearGradient, error) {
	opts := &gradientOptions{}

	for _, opt := range options {
		opt(opts)
	}

	if err := finalizeOptions(opts); err != nil {
		return nil, err
	}

	return &LinearGradient{
		stops:      opts.stops,
		angle:      opts.angle,
		colorspace: opts.colorspace,
	}, nil
}

func (lg *LinearGradient) Colors(steps int) []color.Color {
	return lg.colorspace.Colors(lg.stops, steps)
}

func (lg *LinearGradient) Render(img image.Image) {
	rect := img.Bounds()

	// In linear gradient, angle represents direction,
	// so the actual angle is the opposite.
	theta := 2*math.Pi - lg.angle.NormalizedRadians(rect)
	dx := math.Cos(theta)
	dy := math.Sin(theta)

	corners := [][2]float64{
		{float64(rect.Min.X), float64(rect.Min.Y)},
		{float64(rect.Max.X), float64(rect.Min.Y)},
		{float64(rect.Min.X), float64(rect.Max.Y)},
		{float64(rect.Max.X), float64(rect.Max.Y)},
	}

	minProj := math.MaxFloat64
	maxProj := -math.MaxFloat64

	for _, c := range corners {
		x := c[0]
		y := c[1]

		proj := x*dx + y*dy
		if proj < minProj {
			minProj = proj
		}
		if proj > maxProj {
			maxProj = proj
		}
	}

	length := maxProj - minProj
	colors := lg.Colors(int(math.Round(length)))

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			t := int(math.Round(float64(x)*dx + float64(y)*dy - minProj))
			t = max(0, min(t, len(colors)-1))
			rgb := colors[t].RGB()
			native := gocolor.RGBA{rgb.R, rgb.G, rgb.B, 255}

			img.(draw.Image).Set(x, y, native)
		}
	}
}
