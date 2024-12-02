package color_test

import (
	"fmt"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestHSL(t *testing.T) {
	assert.Implements(t, (*interface {
		CMYK() *color.CMYK
		Hex() string
		HSV() *color.HSV
		RGB() *color.RGB
		XYZ() *color.XYZ

		String() string
	})(nil), color.HSL{})
}

func TestHSL_HSV(t *testing.T) {
	cases := []struct {
		color    *color.HSL
		expected *color.HSV
	}{
		{&color.HSL{0, 100, 50}, &color.HSV{0, 100, 100}},
		{&color.HSL{120, 100, 50}, &color.HSV{120, 100, 100}},
		{&color.HSL{240, 100, 50}, &color.HSV{240, 100, 100}},
		{&color.HSL{10, 16, 79}, &color.HSV{10, 8, 82}},
		{&color.HSL{180, 0, 30}, &color.HSV{180, 0, 30}},
		{&color.HSL{0, 0, 0}, &color.HSV{0, 0, 0}},
		{&color.HSL{30, 20, 60}, &color.HSV{30, 24, 68}},
	}

	for _, c := range cases {
		actual := c.color.HSV()
		assert.Equal(t, c.expected, actual)
	}
}

func ExampleHSL_HSV() {
	c := color.NewHSL(30, 20, 60)
	hsv := c.HSV()

	fmt.Printf("hsv(%d, %d%%, %d%%)\n", hsv.H, hsv.S, hsv.V) // Output: hsv(30, 24%, 68%)
}

func TestHSL_RGB(t *testing.T) {
	cases := []struct {
		color    *color.HSL
		expected *color.RGB
	}{
		{&color.HSL{0, 100, 50}, &color.RGB{255, 0, 0}},
		{&color.HSL{120, 100, 50}, &color.RGB{0, 255, 0}},
		{&color.HSL{240, 100, 50}, &color.RGB{0, 0, 255}},
		{&color.HSL{60, 100, 50}, &color.RGB{255, 255, 0}},
		{&color.HSL{180, 100, 50}, &color.RGB{0, 255, 255}},
		{&color.HSL{300, 100, 50}, &color.RGB{255, 0, 255}},
		{&color.HSL{0, 0, 50}, &color.RGB{128, 128, 128}},
		{&color.HSL{0, 0, 0}, &color.RGB{0, 0, 0}},
		{&color.HSL{0, 0, 100}, &color.RGB{255, 255, 255}},
	}

	for _, c := range cases {
		actual := c.color.RGB()
		assert.Equal(t, c.expected, actual)
	}
}

func ExampleHSL_RGB() {
	c := color.NewHSL(0, 100, 50)
	fmt.Println(c.RGB()) // Output: rgb(255, 0, 0)
}

func TestHSL_Edit(t *testing.T) {
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

func ExampleHSL_Edit() {
	cyan := color.NewHSL(180, 100, 50) // rgb(0, 255, 255)

	// Decrease saturation and print as RGB.
	fmt.Println(cyan.Edit(func(c *color.HSL) {
		c.S = 30
	}).RGB()) // Output: rgb(89, 166, 166)
}
