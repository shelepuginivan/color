package gradient

import (
	"image"
	gocolor "image/color"
	"image/draw"
	"math"

	"github.com/shelepuginivan/color"
)

// ConicGradient is a gradient with a progressive color transition rotating
// around the center.
type ConicGradient struct {
	stops      []*ColorStop
	angle      angleSpec
	center     pointSpec
	colorspace Colorspace
}

// NewConic returns a new instance of [ConicGradient].
//
// Supported options are:
//   - [WithCenterAt] and [WithRelativeCenter] set the rotation axis point.
//   - [WithAngle] and [WithDirection] set the starting angle of rotation.
//     The default angle is 0° (rotation starts at the top).
func NewConic(options ...GradientOption) (*ConicGradient, error) {
	opts := &gradientOptions{}

	for _, opt := range options {
		opt(opts)
	}

	if err := finalizeOptions(opts); err != nil {
		return nil, err
	}

	return &ConicGradient{
		stops:      opts.stops,
		angle:      opts.angle,
		center:     opts.center,
		colorspace: opts.colorspace,
	}, nil
}

func (cg *ConicGradient) Colors(steps int) []color.Color {
	return cg.colorspace.Colors(cg.stops, steps)
}

func (cg *ConicGradient) Render(img image.Image) {
	rect := img.Bounds()
	center := cg.center.Position(rect)
	angle := cg.angle.NormalizedRadiansWithCenter(rect, center)

	depth := 360
	colors := cg.Colors(depth)

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			dx := float64(x - center.x)
			dy := float64(y - center.y)
			angle := math.Atan2(dy, dx) + angle
			angle = math.Mod(angle+2*math.Pi, 2*math.Pi)
			fraction := angle / (2 * math.Pi)
			t := int((fraction*float64(depth))+float64(depth)) % depth

			rgb := colors[t].RGB()
			native := &gocolor.RGBA{rgb.R, rgb.G, rgb.B, 255}

			img.(draw.Image).Set(x, y, native)
		}
	}
}
