package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestLuminance(t *testing.T) {
	cases := []struct {
		color    color.Color
		expected float64
	}{
		{&color.RGB{0, 0, 0}, 0.0},
		{&color.RGB{255, 255, 255}, 1.0},
		{&color.RGB{255, 0, 0}, 0.2126},
		{&color.RGB{0, 255, 0}, 0.7152},
		{&color.RGB{0, 0, 255}, 0.0722},
		{&color.RGB{153, 193, 241}, 0.5126},
		{&color.RGB{255, 163, 72}, 0.4792},
		{&color.RGB{0, 122, 107}, 0.1498},
	}
	for _, c := range cases {
		actual := color.Luminance(c.color)
		assert.InDelta(t, c.expected, actual, 0.0001)
	}
}

func TestContrast(t *testing.T) {
	cases := []struct {
		c1, c2   color.Color
		expected float64
	}{
		{&color.RGB{0, 0, 0}, &color.RGB{0, 0, 0}, 1},
		{&color.RGB{0, 0, 0}, &color.RGB{255, 255, 255}, 21.0},
		{&color.RGB{255, 255, 255}, &color.RGB{0, 0, 0}, 21.0},
		{&color.RGB{0, 0, 255}, &color.RGB{0, 0, 0}, 2.4444},
		{&color.RGB{249, 240, 107}, &color.RGB{165, 29, 45}, 6.2926},
	}
	for _, c := range cases {
		actual := color.Contrast(c.c1, c.c2)
		assert.InDelta(t, c.expected, actual, 0.001)
	}
}

func TestContrastBlackWhite(t *testing.T) {
	cases := []struct {
		color    color.Color
		expected color.Color
	}{
		{&color.RGB{255, 255, 255}, &color.RGB{0, 0, 0}},
		{&color.RGB{0, 0, 0}, &color.RGB{255, 255, 255}},
		{&color.RGB{195, 200, 195}, &color.RGB{0, 0, 0}},
		{&color.RGB{80, 177, 96}, &color.RGB{0, 0, 0}},
		{&color.RGB{100, 149, 237}, &color.RGB{0, 0, 0}},
		{&color.RGB{27, 30, 28}, &color.RGB{255, 255, 255}},
	}
	for _, c := range cases {
		actual := color.ContrastBlackWhite(c.color)
		assert.Equal(t, c.expected, actual)
	}
}
