package gradient

import "github.com/shelepuginivan/color"

// Colorspace defines how the intermediate colors of the gradient are
// calculated.
type Colorspace interface {
	// Colors calculates intermediate colors between stops for the specified
	// number of steps.
	Colors(stops []*ColorStop, steps int) []color.Color

	// Intermediate calculates intermediate colors between start and end.
	Intermediate(start, end color.Color, steps int) []color.Color
}
