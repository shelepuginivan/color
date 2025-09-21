package gradient

import (
	"math"

	"github.com/shelepuginivan/color"
	"github.com/shelepuginivan/color/internal/interpolate"
)

type ColorspaceOklch struct {
	method HueInterpolationMethod
}

func (cOklch *ColorspaceOklch) Colors(stops []*ColorStop, steps int) []color.Color {
	colors := make([]color.Color, 0, steps)

	for stopIndex := range len(stops) - 1 {
		first := stops[stopIndex]
		second := stops[stopIndex+1]

		stepFraction := second.Position - first.Position
		segmentSteps := int(math.Round(float64(steps) * stepFraction))

		start := first.Color.RGB()
		end := second.Color.RGB()

		colors = append(colors, cOklch.Intermediate(start, end, segmentSteps)...)
	}

	return colors
}

func (cOklch *ColorspaceOklch) Intermediate(start, end color.Color, steps int) []color.Color {
	var (
		colors = make([]color.Color, steps)

		s = start.Oklch()
		e = end.Oklch()
	)

	direction, angle := interpolate.Hue(s.H, e.H, interpolate.HueInterpolationMethod(cOklch.method))

	currentHue := float64(s.H)
	dHue := float64(direction) * float64(angle) / float64(steps-1)

	for i := range steps {
		scale := float64(i) / float64(steps-1)

		hue := int(math.Round(currentHue))
		l := interpolate.Rectangular(s.L, e.L, scale)
		c := interpolate.Rectangular(s.C, e.C, scale)

		colors[i] = &color.Oklch{
			L: l,
			C: c,
			H: hue,
		}

		currentHue = math.Mod(currentHue+dHue, 360)
	}

	return colors
}
