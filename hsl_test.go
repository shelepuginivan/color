package color_test

import (
	"fmt"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

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
	rgb := c.RGB()

	fmt.Printf("rgb(%d, %d, %d)\n", rgb.R, rgb.G, rgb.B) // Output: rgb(255, 0, 0)
}
