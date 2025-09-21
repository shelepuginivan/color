package interpolate

type HueInterpolationMethod uint8

const (
	ShorterHue HueInterpolationMethod = iota
	LongerHue
	IncreasingHue
	DecreasingHue
)

// Hue calculates direction and angle of hue interpolation on color wheel.
func Hue(start, end int, method HueInterpolationMethod) (direction int, angle int) {
	// SEE: https://developer.mozilla.org/en-US/docs/Web/CSS/hue-interpolation-method
	switch method {
	case ShorterHue:
		direction, angle = calcShorterHue(start, end)
	case LongerHue:
		direction, angle = calcLongerHue(start, end)
	case IncreasingHue:
		direction, angle = calcIncreasingHue(start, end)
	case DecreasingHue:
		direction, angle = calcDecreasingHue(start, end)
	}

	return
}

func calcShorterHue(start, end int) (direction int, angle int) {
	diff := (end - start + 360) % 360

	if diff <= 180 {
		direction = 1
		angle = diff
	} else {
		direction = -1
		angle = 360 - diff
	}

	return
}

func calcLongerHue(start, end int) (direction int, angle int) {
	diff := (end - start + 360) % 360

	if diff <= 180 {
		direction = -1
		angle = 360 - diff
	} else {
		direction = 1
		angle = diff
	}

	return
}

func calcIncreasingHue(start, end int) (direction int, angle int) {
	direction = 1
	angle = (end - start + 360) % 360
	return
}

func calcDecreasingHue(start, end int) (direction int, angle int) {
	direction = -1
	angle = (start - end + 360) % 360
	return
}
