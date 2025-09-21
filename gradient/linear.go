package gradient

import "github.com/shelepuginivan/color"

type LinearGradient struct {
	stops      []*ColorStop
	colorspace Colorspace
}

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
		colorspace: opts.colorspace,
	}, nil
}

func (lg *LinearGradient) Colors(steps int) []color.Color {
	return lg.colorspace.Colors(lg.stops, steps)
}
