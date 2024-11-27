package color_test

import (
	"fmt"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

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
