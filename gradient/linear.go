package gradient

import "github.com/shelepuginivan/color"

type LinearGradient struct {
	start color.Color
	end   color.Color
}

func NewLinear(start, end color.Color) *LinearGradient {
	return &LinearGradient{start, end}
}

func (lg *LinearGradient) Colors(steps int) []color.Color {
	var (
		colors = make([]color.Color, steps)
		start  = lg.start.RGB()
		end    = lg.end.RGB()
	)

	for i := range steps {
		scale := float64(i) / float64(steps)

		r := float64(start.R)*(1-scale) + float64(end.R)*scale
		g := float64(start.G)*(1-scale) + float64(end.G)*scale
		b := float64(start.B)*(1-scale) + float64(end.B)*scale

		colors[i] = &color.RGB{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
		}
	}

	return colors
}
