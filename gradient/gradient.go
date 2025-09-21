package gradient

import (
	"image"

	"github.com/shelepuginivan/color"
)

type Gradient interface {
	Colors(steps int) []color.Color
	Render(img image.Image)
}
