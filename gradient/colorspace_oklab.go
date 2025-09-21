package gradient

import (
	"math"

	"github.com/shelepuginivan/color"
	"github.com/shelepuginivan/color/internal/interpolate"
)

type ColorspaceOklab struct{}

func (cOklab *ColorspaceOklab) Colors(stops []*ColorStop, steps int) []color.Color {
	colors := make([]color.Color, 0, steps)

	for stopIndex := range len(stops) - 1 {
		first := stops[stopIndex]
		second := stops[stopIndex+1]

		stepFraction := second.Position - first.Position
		segmentSteps := int(math.Round(float64(steps) * stepFraction))

		colors = append(colors, cOklab.Intermediate(
			first.Color,
			second.Color,
			segmentSteps,
		)...)
	}

	return colors
}

func (cOklab *ColorspaceOklab) Intermediate(start, end color.Color, steps int) []color.Color {
	var (
		colors = make([]color.Color, steps)
		s      = start.Oklab()
		e      = end.Oklab()
	)

	for i := range steps {
		scale := float64(i) / float64(steps-1)

		l := interpolate.Rectangular(s.L, e.L, scale)
		a := interpolate.Rectangular(s.A, e.A, scale)
		b := interpolate.Rectangular(s.B, e.B, scale)

		colors[i] = &color.Oklab{
			L: l,
			A: a,
			B: b,
		}
	}

	return colors
}
