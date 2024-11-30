package color_test

import (
	"fmt"
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
		{&color.HSV{300, 100, 50}, &color.RGB{128, 0, 128}},
	}

	for _, c := range tests {
		actual := c.color.RGB()
		assert.Equal(t, c.expected, actual)
	}
}

func TestHSV_Edit(t *testing.T) {
	cases := []struct {
		color    *color.HSL
		expected *color.HSL
		editfn   func(*color.HSL)
	}{
		{
			color:    &color.HSL{H: 120, S: 100, L: 50},
			expected: &color.HSL{H: 120, S: 50, L: 50},
			editfn: func(c *color.HSL) {
				c.S = 50
			},
		},
		{
			color:    &color.HSL{H: 240, S: 100, L: 50},
			expected: &color.HSL{H: 240, S: 100, L: 75},
			editfn: func(c *color.HSL) {
				c.L = 75
			},
		},
		{
			color:    &color.HSL{H: 0, S: 0, L: 0},
			expected: &color.HSL{H: 0, S: 0, L: 100},
			editfn: func(c *color.HSL) {
				c.L = 100
			},
		},
		{
			color:    &color.HSL{H: 360, S: 100, L: 100},
			expected: &color.HSL{H: 360, S: 100, L: 0},
			editfn: func(c *color.HSL) {
				c.L = 0
			},
		},
	}

	for _, c := range cases {
		actual := c.color.Edit(c.editfn)
		assert.Equal(t, c.expected, actual)
		assert.Equal(t, c.expected, c.color)
	}
}

func ExampleHSV_Edit() {
	green := color.NewHSV(72, 100, 40) // rgb(78, 102, 0)

	// Increase value and print as RGB.
	fmt.Println(green.Edit(func(c *color.HSV) {
		c.V = 100
	}).RGB()) // Output: rgb(204, 255, 0)
}
