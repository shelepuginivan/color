package gradient

import (
	"math"

	"github.com/shelepuginivan/color"
	"github.com/shelepuginivan/color/internal/interpolate"
)

type ColorspaceHSL struct {
	hueType HueType
}

func (cHSL *ColorspaceHSL) Colors(stops []*ColorStop, steps int) []color.Color {
	colors := make([]color.Color, 0, steps)

	for stopIndex := range len(stops) - 1 {
		first := stops[stopIndex]
		second := stops[stopIndex+1]

		stepFraction := second.Position - first.Position
		segmentSteps := int(math.Round(float64(steps) * stepFraction))

		start := first.Color.RGB()
		end := second.Color.RGB()

		colors = append(colors, cHSL.Intermediate(start, end, segmentSteps)...)
	}

	return colors
}

func (cHSL *ColorspaceHSL) Intermediate(start, end color.Color, steps int) []color.Color {
	// NOTE: see [Hue Interpolation] on MDN.
	//
	// [Hue Interpolation]: https://developer.mozilla.org/en-US/docs/Web/CSS/hue-interpolation-method
	var (
		colors = make([]color.Color, steps)

		s = start.HSL()
		e = end.HSL()
	)

	direction, angle := interpolate.Hue(s.H, e.H, interpolate.HueInterpolationMethod(cHSL.hueType))
	currentHue := float64(s.H)
	dHue := float64(direction) * float64(angle) / float64(steps-1)

	for i := range steps {
		scale := float64(i) / float64(steps-1)

		hue := int(math.Round(currentHue))
		saturation := interpolate.RectangularInt(s.S, e.S, scale)
		lightness := interpolate.RectangularInt(s.L, e.L, scale)

		colors[i] = &color.HSL{
			H: hue,
			S: saturation,
			L: lightness,
		}

		currentHue = math.Mod(currentHue+dHue, 360)
	}

	return colors
}
