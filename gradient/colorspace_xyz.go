package gradient

import (
	"math"

	"github.com/shelepuginivan/color"
)

type ColorspaceXYZ struct {
	whitepoint *color.XYZ
}

func (cXYZ *ColorspaceXYZ) Colors(stops []*ColorStop, steps int) []color.Color {
	colors := make([]color.Color, 0, steps)

	for stopIndex := range len(stops) - 1 {
		first := stops[stopIndex]
		second := stops[stopIndex+1]

		stepFraction := second.Position - first.Position
		segmentSteps := int(math.Round(float64(steps) * stepFraction))

		colors = append(colors, cXYZ.Intermediate(
			first.Color,
			second.Color,
			segmentSteps,
		)...)
	}

	return colors
}

func (cXYZ *ColorspaceXYZ) Intermediate(start, end color.Color, steps int) []color.Color {
	var (
		colors = make([]color.Color, steps)
		s      = start.Lab().XYZWithWhitepoint(cXYZ.whitepoint)
		e      = end.Lab().XYZWithWhitepoint(cXYZ.whitepoint)
	)

	for i := range steps {
		scale := float64(i) / float64(steps-1)

		x := s.X*(1-scale) + e.X*scale
		y := s.Y*(1-scale) + e.Y*scale
		z := s.Z*(1-scale) + e.Z*scale

		colors[i] = &color.XYZ{
			X: x,
			Y: y,
			Z: z,
		}
	}

	return colors
}
