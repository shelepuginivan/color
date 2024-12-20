package color_test

import (
	"math/rand"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestShades(t *testing.T) {
	for range 1000 {
		var (
			h = rand.Intn(360)
			s = rand.Intn(101)
			l = rand.Intn(101)

			n = 2 + rand.Intn(29)
			c = &color.HSL{h, s, l}
		)

		shades := color.Shades(c, n)

		assert.Equal(t, c.HSL(), shades[0])
		assert.Equal(t, &color.RGB{0, 0, 0}, shades[n-1].RGB())

		contrastToBlack := make([]float64, n)

		for i, shade := range shades {
			contrastToBlack[i] = color.Contrast(shade, &color.RGB{0, 0, 0})
		}

		assert.IsNonIncreasing(t, contrastToBlack)
	}
}

func TestTints(t *testing.T) {
	for range 1000 {
		var (
			h = rand.Intn(360)
			s = rand.Intn(101)
			l = rand.Intn(101)

			n = 2 + rand.Intn(29)
			c = &color.HSL{h, s, l}
		)

		tints := color.Tints(c, n)

		assert.Equal(t, c.HSL(), tints[0])
		assert.Equal(t, &color.RGB{255, 255, 255}, tints[n-1].RGB())

		contrastToWhite := make([]float64, n)

		for i, tint := range tints {
			contrastToWhite[i] = color.Contrast(tint, &color.RGB{255, 255, 255})
		}

		assert.IsNonIncreasing(t, contrastToWhite)
	}
}

func TestTones(t *testing.T) {
	for range 1000 {
		var (
			h = rand.Intn(360)
			s = rand.Intn(101)
			l = rand.Intn(101)

			n = 2 + rand.Intn(29)
			c = &color.HSL{h, s, l}
		)

		tones := color.Tones(c, n)

		assert.Equal(t, c.HSL(), tones[0])
		assert.Equal(t, &color.RGB{128, 128, 128}, tones[n-1].RGB())
	}
}
