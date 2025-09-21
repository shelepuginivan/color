package gradient

import (
	"math"

	"github.com/shelepuginivan/color"
)

type ColorspaceRGB struct{}

func (cRGB *ColorspaceRGB) Colors(stops []*ColorStop, steps int) []color.Color {
	colors := make([]color.Color, 0, steps)

	for stopIndex := range len(stops) - 1 {
		first := stops[stopIndex]
		second := stops[stopIndex+1]

		stepFraction := second.Position - first.Position
		segmentSteps := int(math.Round(float64(steps) * stepFraction))

		start := first.Color.RGB()
		end := second.Color.RGB()

		for i := range segmentSteps {
			scale := float64(i) / float64(segmentSteps)

			r := float64(start.R)*(1-scale) + float64(end.R)*scale
			g := float64(start.G)*(1-scale) + float64(end.G)*scale
			b := float64(start.B)*(1-scale) + float64(end.B)*scale

			colors = append(colors, &color.RGB{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
			})
		}
	}

	return colors

}
