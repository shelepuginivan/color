package color

import "math"

// Shades returns n shades of the color. New colors are created by reducing
// lightness of the original color, making them darker.
//
// The first element of the resulting slice is the color itself. The last is
// black. Special case is n < 2: a slice with a single element being the color
// itself is returned.
func Shades(color Color, n int) []Color {
	hsl := color.HSL()

	if n < 2 {
		return []Color{hsl}
	}

	shades := make([]Color, n)
	shades[0] = hsl

	floatL := float64(hsl.L)
	step := floatL / float64(n-1)

	for i := 1; i < n; i++ {
		decrement := step * float64(i)
		newL := int(math.Round(floatL - decrement))
		normalizedL := max(newL, 0)

		shades[i] = &HSL{
			H: hsl.H,
			S: hsl.S,
			L: normalizedL,
		}
	}

	return shades
}

// Tints returns n tints shades of the color. New colors are created by
// increasing lightness of the original color, making them lighter.
//
// The first element of the resulting slice is the color itself. The last is
// white. Special case is n < 2: a slice with a single element being the color
// itself is returned.
func Tints(color Color, n int) []Color {
	hsl := color.HSL()

	if n < 2 {
		return []Color{hsl}
	}

	tints := make([]Color, n)
	tints[0] = hsl

	floatL := float64(hsl.L)
	step := (100 - floatL) / float64(n-1)

	for i := 1; i < n; i++ {
		increment := step * float64(i)
		newL := int(math.Round(floatL + increment))
		normalizedL := min(newL, 100)

		tints[i] = &HSL{
			H: hsl.H,
			S: hsl.S,
			L: normalizedL,
		}
	}

	return tints
}

// Tones returns n tones of the color. New colors are created by reducing
// saturation of the original color and changing lightness towards 50% (in
// HSL), making them more grayish.
//
// The first element of the resulting slice is the color itself. The last is
// gray. Special case is n < 2: a slice with a single element being the color
// itself is returned.
func Tones(color Color, n int) []Color {
	hsl := color.HSL()

	if n < 2 {
		return []Color{hsl}
	}

	tones := make([]Color, n)
	tones[0] = hsl

	floatL := float64(hsl.L)
	floatS := float64(hsl.S)
	floatCount := float64(n - 1)

	stepS := floatS / floatCount
	stepL := (50 - floatL) / floatCount

	for i := 1; i < n; i++ {
		var normalizedL int

		// Normalize new lightness. If the original color is lighter than gray,
		// decrease the lightness, otherwise increase it. Ensure that lightness
		// is not lower / greater than gray respectively.
		if hsl.L > 50 {
			decrement := stepL * float64(i)
			newL := int(math.Round(floatL - decrement))
			normalizedL = min(newL, 50)
		} else {
			increment := stepL * float64(i)
			newL := int(math.Round(floatL + increment))
			normalizedL = max(newL, 50)
		}

		decrementS := stepS * float64(i)
		newS := int(math.Round(floatS - decrementS))
		normalizedS := max(newS, 0)

		tones[i] = &HSL{
			H: hsl.H,
			S: normalizedS,
			L: normalizedL,
		}
	}

	return tones
}
