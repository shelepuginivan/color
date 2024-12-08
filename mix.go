package color

import (
	"math"

	"github.com/shelepuginivan/color/internal/degrees"
	"github.com/shelepuginivan/color/internal/percents"
)

// MixCMYK calculates the average color in CMYK colorspace from an arbitrary
// number of colors.
func MixCMYK(colors ...Color) *CMYK {
	if len(colors) == 0 {
		return &CMYK{0, 0, 0, 0}
	}

	var (
		c, m, y, k float64
		total      = float64(len(colors))
	)

	for _, color := range colors {
		cmyk := color.CMYK()

		c += percents.ToFloat(cmyk.C)
		m += percents.ToFloat(cmyk.M)
		y += percents.ToFloat(cmyk.Y)
		k += percents.ToFloat(cmyk.K)
	}

	return &CMYK{
		C: percents.FromFloat(c / total),
		M: percents.FromFloat(m / total),
		Y: percents.FromFloat(y / total),
		K: percents.FromFloat(k / total),
	}
}

// MixHSL calculates the average color in HSL colorspace from an arbitrary
// number of colors.
func MixHSL(colors ...Color) *HSL {
	if len(colors) == 0 {
		return &HSL{0, 0, 0}
	}

	var (
		sumX = 0.0
		sumY = 0.0
		sumS = 0.0
		sumL = 0.0
		n    = float64(len(colors))
	)

	for _, color := range colors {
		hsl := color.HSL()

		// Convert hue to radians
		rad := degrees.ToRadians(hsl.H)

		// Convert hue to a point on a circle
		sumX += math.Cos(rad)
		sumY += math.Sin(rad)

		// Sum up saturation and lightness
		sumS += float64(hsl.S)
		sumL += float64(hsl.L)
	}

	// Average the points on the circle and convert back to hue.
	var (
		arctan = math.Atan2(sumY/n, sumX/n)
		avgH   = degrees.FromRadians(arctan)
	)

	// Average saturation and lightness.
	var (
		avgS = int(math.Round(sumS / n))
		avgL = int(math.Round(sumL / n))
	)

	return &HSL{avgH, avgS, avgL}
}

// MixLab calculates the average color in Lab colorspace from an arbitrary
// number of colors.
func MixLab(colors ...Color) *Lab {
	if len(colors) == 0 {
		return &Lab{0, 0, 0}
	}

	var (
		sumL  = 0.0
		sumA  = 0.0
		sumB  = 0.0
		total = float64(len(colors))
	)

	for _, color := range colors {
		lab := color.Lab()

		sumL += lab.L
		sumA += lab.A
		sumB += lab.B
	}

	var (
		avgL = sumL / total
		avgA = sumA / total
		avgB = sumB / total
	)

	return &Lab{avgL, avgA, avgB}
}

// MixRGB calculates the average color in RGB colorspace from an arbitrary
// number of colors.
func MixRGB(colors ...Color) *RGB {
	if len(colors) == 0 {
		return &RGB{0, 0, 0}
	}

	var (
		r, g, b float64
		total   = float64(len(colors))
	)

	for _, color := range colors {
		rgb := color.RGB()
		r += float64(rgb.R)
		g += float64(rgb.G)
		b += float64(rgb.B)
	}

	return &RGB{
		R: uint8(math.Round(r / total)),
		G: uint8(math.Round(g / total)),
		B: uint8(math.Round(b / total)),
	}
}
