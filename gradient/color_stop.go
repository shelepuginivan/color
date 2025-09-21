package gradient

import "github.com/shelepuginivan/color"

// ColorStop represents a point in a gradient where a specific color is defined
// at a certain position.
type ColorStop struct {
	Color    color.Color
	Position float64
}
