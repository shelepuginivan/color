package gradient

import (
	"fmt"
	"sort"

	"github.com/shelepuginivan/color"
)

// HueInterpolationMethod represents how the hue should transition between
// colors along the hue circle. Used for gradients in cylindrical color spaces.
type HueInterpolationMethod uint8

const (
	// Follow the shortest path on the hue circle between colors
	// (minimal angle difference).
	ShorterHue HueInterpolationMethod = iota

	// Follow the longer path on the hue circle between colors
	// (maximal angle difference).
	LongerHue

	// Always increase hue angle from start to end, moving forward around the
	// circle.
	IncreasingHue

	// Always decrease hue angle from start to end, moving backward around the
	// circle.
	DecreasingHue
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
	position = max(0, min(1, position))

	return func(opts *gradientOptions) {
		opts.stops = append(opts.stops, &ColorStop{color, position})
	}
}

// InRGB sets gradient colorspace to RGB.
func InRGB(opts *gradientOptions) {
	opts.colorspace = &ColorspaceRGB{}
}

// InHSL sets gradient colorspace to HSL. The method parameter controls how the
// hue should transition between colors along the hue circle.
// See [HueInterpolationMethod] for more information.
func InHSL(method HueInterpolationMethod) GradientOption {
	return func(opts *gradientOptions) {
		opts.colorspace = &ColorspaceHSL{method}
	}
}

// InHSV sets gradient colorspace to HSV. The method parameter controls how the
// hue should transition between colors along the hue circle.
// See [HueInterpolationMethod] for more information.
func InHSV(method HueInterpolationMethod) GradientOption {
	return func(opts *gradientOptions) {
		opts.colorspace = &ColorspaceHSV{method}
	}
}

// InXYZ sets gradient colorspace to CIE XYZ. The whitepoint parameter
// specifies reference white color.
func InXYZ(whitepoint *color.XYZ) GradientOption {
	if whitepoint == nil {
		whitepoint = color.D65
	}

	return func(opts *gradientOptions) {
		opts.colorspace = &ColorspaceXYZ{whitepoint}
	}
}

// InLab sets gradient colorspace to CIE Lab. The whitepoint parameter
// specifies reference white color.
func InLab(whitepoint *color.XYZ) GradientOption {
	if whitepoint == nil {
		whitepoint = color.D65
	}

	return func(opts *gradientOptions) {
		opts.colorspace = &ColorspaceLab{whitepoint}
	}
}

// InLch sets gradient colorspace to Lch, a cylindrical counterpart of CIE Lab.
//
// The method parameter controls how the hue should transition between colors
// along the hue circle. See [HueInterpolationMethod] for more information.
//
// The whitepoint parameter specifies reference white color.
func InLch(method HueInterpolationMethod, whitepoint *color.XYZ) GradientOption {
	if whitepoint == nil {
		whitepoint = color.D65
	}

	return func(opts *gradientOptions) {
		opts.colorspace = &ColorspaceLch{method, whitepoint}
	}
}

// InOklab sets gradient colorspace to Oklab.
func InOklab(opts *gradientOptions) {
	opts.colorspace = &ColorspaceOklab{}
}

// InOklch sets gradient colorspace to Oklch, a cylindrical counterpart of Oklab.
// The method parameter controls how the hue should transition between colors
// along the hue circle. See [HueInterpolationMethod] for more information.
func InOklch(method HueInterpolationMethod) GradientOption {
	return func(opts *gradientOptions) {
		opts.colorspace = &ColorspaceOklch{method}
	}
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
