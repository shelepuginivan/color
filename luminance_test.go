package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestLuminance(t *testing.T) {
	cases := []struct {
		color    *color.Color
		expected float64
	}{
		{&color.Color{0, 0, 0}, 0.0},
		{&color.Color{255, 255, 255}, 1.0},
		{&color.Color{255, 0, 0}, color.LuminanceRed},
		{&color.Color{0, 255, 0}, color.LuminanceGreen},
		{&color.Color{0, 0, 255}, color.LuminanceBlue},
		{&color.Color{153, 193, 241}, 0.5126},
		{&color.Color{255, 163, 72}, 0.4792},
		{&color.Color{0, 122, 107}, 0.1498},
	}

	for _, c := range cases {
		actual := color.Luminance(c.color)
		assert.InDelta(t, c.expected, actual, 0.0001)
	}
}

func TestContrast(t *testing.T) {
	cases := []struct {
		c1, c2   *color.Color
		expected float64
	}{
		{&color.Color{0, 0, 0}, &color.Color{0, 0, 0}, 1},
		{&color.Color{0, 0, 0}, &color.Color{255, 255, 255}, 21.0},
		{&color.Color{255, 255, 255}, &color.Color{0, 0, 0}, 21.0},
		{&color.Color{0, 0, 255}, &color.Color{0, 0, 0}, 2.4444},
		{&color.Color{249, 240, 107}, &color.Color{165, 29, 45}, 6.2926},
	}

	for _, c := range cases {
		actual := color.Contrast(c.c1, c.c2)
		assert.InDelta(t, c.expected, actual, 0.001)
	}
}

func TestContrastBlackWhite(t *testing.T) {
	cases := []struct {
		color    *color.Color
		expected *color.Color
	}{
		{&color.Color{255, 255, 255}, &color.Color{0, 0, 0}},
		{&color.Color{0, 0, 0}, &color.Color{255, 255, 255}},
		{&color.Color{195, 200, 195}, &color.Color{0, 0, 0}},
		{&color.Color{80, 177, 96}, &color.Color{0, 0, 0}},
		{&color.Color{100, 149, 237}, &color.Color{0, 0, 0}},
		{&color.Color{27, 30, 28}, &color.Color{255, 255, 255}},
	}

	for _, c := range cases {
		actual := color.ContrastBlackWhite(c.color)
		assert.Equal(t, c.expected, actual)
	}
}
