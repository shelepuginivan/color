package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestNewHSV(t *testing.T) {
	tests := []struct {
		h, s, v  int
		expected *color.HSV
	}{
		{0, 100, 100, &color.HSV{H: 0, S: 100, V: 100}},
		{360, 50, 50, &color.HSV{H: 0, S: 50, V: 50}},
		{180, 200, 150, &color.HSV{H: 180, S: 100, V: 100}},
	}

	for _, c := range tests {
		actual := color.NewHSV(c.h, c.s, c.v)
		assert.Equal(t, c.expected, actual)
	}
}

func TestHSV_HSL(t *testing.T) {
	tests := []struct {
		color    *color.HSV
		expected *color.HSL
	}{
		{&color.HSV{H: 0, S: 100, V: 100}, &color.HSL{H: 0, S: 100, L: 50}},
		{&color.HSV{H: 120, S: 100, V: 100}, &color.HSL{H: 120, S: 100, L: 50}},
		{&color.HSV{H: 240, S: 100, V: 100}, &color.HSL{H: 240, S: 100, L: 50}},
		{&color.HSV{H: 0, S: 0, V: 0}, &color.HSL{H: 0, S: 0, L: 0}},
		{&color.HSV{H: 0, S: 0, V: 100}, &color.HSL{H: 0, S: 0, L: 100}},
	}

	for _, c := range tests {
		actual := c.color.HSL()
		assert.Equal(t, c.expected, actual)
	}
}

func TestHSV_RGB(t *testing.T) {
	tests := []struct {
		color    *color.HSV
		expected *color.RGB
	}{
		{&color.HSV{0, 100, 100}, &color.RGB{255, 0, 0}},
		{&color.HSV{120, 100, 100}, &color.RGB{0, 255, 0}},
		{&color.HSV{240, 100, 100}, &color.RGB{0, 0, 255}},
	}

	for _, c := range tests {
		actual := c.color.RGB()
		assert.Equal(t, c.expected, actual)
	}
}
