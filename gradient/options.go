package gradient

import (
	"fmt"
	"sort"

	"github.com/shelepuginivan/color"
)

type gradientOptions struct {
	stops      []*ColorStop
	colorspace Colorspace
}

// GradientOption is a functional options type for configuring [Gradient].
type GradientOption func(*gradientOptions)

// WithColorStop appends a color stop to the gradient.
//
// The position should be a number within [0, 1], otherwise it is clamped.
func WithColorStop(color color.Color, position float64) GradientOption {
	return func(opts *gradientOptions) {
		opts.stops = append(opts.stops, &ColorStop{color, position})
	}
}

// InRGB sets gradient colorspace to RGB.
func InRGB(opts *gradientOptions) {
	opts.colorspace = &ColorspaceRGB{}
}

// finalizeOptions validates and normalizes gradient options.
func finalizeOptions(opts *gradientOptions) error {
	if len(opts.stops) < 2 {
		return fmt.Errorf("gradient must contain at least two color stops")
	}

	// Ensure color stops are properly ordered.
	sort.Slice(opts.stops, func(i, j int) bool {
		return opts.stops[i].Position < opts.stops[j].Position
	})

	// Fill color stop at the gradient beginning if it is missing.
	if first := opts.stops[0]; first.Position > 0 {
		opts.stops = append([]*ColorStop{{first.Color, 0}}, opts.stops...)
	}

	// Fill color stop at the gradient end if it is missing.
	if last := opts.stops[len(opts.stops)-1]; last.Position < 1 {
		opts.stops = append(opts.stops, &ColorStop{last.Color, 1})
	}

	// The default colorspace is RGB.
	if opts.colorspace == nil {
		opts.colorspace = &ColorspaceRGB{}
	}

	return nil
}
