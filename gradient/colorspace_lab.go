package gradient

import (
	"math"

	"github.com/shelepuginivan/color"
)

type ColorspaceLab struct {
	whitepoint *color.XYZ
}

func (cLab *ColorspaceLab) Colors(stops []*ColorStop, steps int) []color.Color {
	colors := make([]color.Color, 0, steps)

	for stopIndex := range len(stops) - 1 {
		first := stops[stopIndex]
		second := stops[stopIndex+1]

		stepFraction := second.Position - first.Position
		segmentSteps := int(math.Round(float64(steps) * stepFraction))

		colors = append(colors, cLab.Intermediate(
			first.Color,
			second.Color,
			segmentSteps,
		)...)
	}

	return colors
}

func (cLab *ColorspaceLab) Intermediate(start, end color.Color, steps int) []color.Color {
	var (
		colors = make([]color.Color, steps)
		s      = start.XYZ().LabWithWhitepoint(cLab.whitepoint)
		e      = end.XYZ().LabWithWhitepoint(cLab.whitepoint)
	)

	for i := range steps {
		scale := float64(i) / float64(steps-1)

		l := s.L*(1-scale) + e.L*scale
		a := s.A*(1-scale) + e.A*scale
		b := s.B*(1-scale) + e.B*scale

		colors[i] = &color.Lab{
			L: l,
			A: a,
			B: b,
		}
	}

	return colors
}
