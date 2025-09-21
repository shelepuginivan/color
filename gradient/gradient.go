// Package gradient provides methods for creating color gradients.
package gradient

import (
	"image"

	"github.com/shelepuginivan/color"
)

// Gradient is a common interface for different gradient types.
type Gradient interface {
	// Colors calculates the given number of intermediate colors.
	Colors(steps int) []color.Color

	// Render renders gradient onto image.
	Render(img image.Image)
}
