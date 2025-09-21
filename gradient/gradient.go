package gradient

import "github.com/shelepuginivan/color"

type Gradient interface {
	Colors(steps int) []color.Color
}
