package color

import (
	"math"

	"github.com/shelepuginivan/color/internal/normalize"
)

// MixCMYK calculates the average color in CMYK colorspace from an arbitrary
// number of colors.
func MixCMYK(colors ...*CMYK) *CMYK {
	if len(colors) == 0 {
		return &CMYK{0, 0, 0, 0}
	}

	var (
		c, m, y, k float64
		total      = float64(len(colors))
	)

	for _, color := range colors {
		c += normalize.PercentsFloat(color.C)
		m += normalize.PercentsFloat(color.M)
		y += normalize.PercentsFloat(color.Y)
		k += normalize.PercentsFloat(color.K)
	}

	return &CMYK{
		C: normalize.FloatPercents(c / total),
		M: normalize.FloatPercents(m / total),
		Y: normalize.FloatPercents(y / total),
		K: normalize.FloatPercents(k / total),
	}
}

// MixHSL calculates the average color in HSL colorspace from an arbitrary
// number of colors.
func MixHSL(colors ...*HSL) *HSL {
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
		// Convert hue to radians
		rad := float64(color.H) * math.Pi / 180.0

		// Convert hue to a point on a circle
		sumX += math.Cos(rad)
		sumY += math.Sin(rad)

		// Sum up saturation and lightness
		sumS += float64(color.S)
		sumL += float64(color.L)
	}

	// Average the points on the circle and convert back to hue
	hue := math.Atan2(sumY/n, sumX/n) * 180.0 / math.Pi

	// Average saturation and lightness
	avgH := normalize.Degrees(int(math.Round(hue)))
	avgS := int(math.Round(sumS / n))
	avgL := int(math.Round(sumL / n))

	return &HSL{H: avgH, S: avgS, L: avgL}
}

// MixLab calculates the average color in Lab colorspace from an arbitrary
// number of colors.
//
// (95.047, 100.000, 108.883) is set as the reference white of the resulting
// color. Use [Lab.SetReferenceWhite] to set a different reference white.
func MixLab(colors ...*Lab) *Lab {
	if len(colors) == 0 {
		return &Lab{0, 0, 0, DefaultReferenceWhite}
	}

	var (
		sumL  = 0.0
		sumA  = 0.0
		sumB  = 0.0
		total = float64(len(colors))
	)

	for _, color := range colors {
		norm := &Lab{color.L, color.A, color.B, color.White}
		norm.SetReferenceWhite(DefaultReferenceWhite)

		sumL += norm.L
		sumA += norm.A
		sumB += norm.B
	}

	var (
		avgL = sumL / total
		avgA = sumA / total
		avgB = sumB / total
	)

	return &Lab{avgL, avgA, avgB, DefaultReferenceWhite}
}

// MixRGB calculates the average color in RGB colorspace from an arbitrary
// number of colors.
func MixRGB(colors ...*RGB) *RGB {
	if len(colors) == 0 {
		return &RGB{0, 0, 0}
	}

	var (
		r, g, b float64
		total   = float64(len(colors))
	)

	for _, color := range colors {
		r += float64(color.R)
		g += float64(color.G)
		b += float64(color.B)
	}

	return &RGB{
		R: uint8(math.Round(r / total)),
		G: uint8(math.Round(g / total)),
		B: uint8(math.Round(b / total)),
	}
}
