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

		// Direction of the hue transition.
		// 1 for clockwise, -1 for counter-clockwise.
		direction int

		// Hue difference angle.
		angle int
	)

	switch cHSL.hueType {
	case ShorterHue:
		direction, angle = cHSL.calcShorterHue(s, e)
	case LongerHue:
		direction, angle = cHSL.calcLongerHue(s, e)
	case IncreasingHue:
		direction, angle = cHSL.calcIncreasingHue(s, e)
	case DecreasingHue:
		direction, angle = cHSL.calcDecreasingHue(s, e)
	}

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

func (cHSL *ColorspaceHSL) calcShorterHue(start, end *color.HSL) (direction int, angle int) {
	diff := (end.H - start.H + 360) % 360

	if diff <= 180 {
		direction = 1
		angle = diff
	} else {
		direction = -1
		angle = 360 - diff
	}

	return
}

func (cHSL *ColorspaceHSL) calcLongerHue(start, end *color.HSL) (direction int, angle int) {
	diff := (end.H - start.H + 360) % 360

	if diff <= 180 {
		direction = -1
		angle = 360 - diff
	} else {
		direction = 1
		angle = diff
	}

	return
}

func (cHSL *ColorspaceHSL) calcIncreasingHue(start, end *color.HSL) (direction int, angle int) {
	direction = 1
	angle = (end.H - start.H + 360) % 360
	return
}

func (cHSL *ColorspaceHSL) calcDecreasingHue(start, end *color.HSL) (direction int, angle int) {
	direction = -1
	angle = (start.H - end.H + 360) % 360
	return
}
