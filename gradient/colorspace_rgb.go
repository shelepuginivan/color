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

		colors = append(colors, cRGB.Intermediate(
			first.Color,
			second.Color,
			segmentSteps,
		)...)
	}

	return colors
}

func (cRGB *ColorspaceRGB) Intermediate(start, end color.Color, steps int) []color.Color {
	var (
		colors = make([]color.Color, steps)
		s      = start.RGB()
		e      = end.RGB()
	)

	for i := range steps {
		scale := float64(i) / float64(steps-1)

		r := float64(s.R)*(1-scale) + float64(e.R)*scale
		g := float64(s.G)*(1-scale) + float64(e.G)*scale
		b := float64(s.B)*(1-scale) + float64(e.B)*scale

		colors[i] = &color.RGB{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
		}
	}

	return colors
}
