package gradient

import "github.com/shelepuginivan/color"

// ColorStop represents a point in a gradient where a specific color is defined
// at a certain position.
//
// Position is a floating-point number in range [0, 1].
type ColorStop struct {
	Color    color.Color
	Position float64
}
