package color_test

import (
	"fmt"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestMixCMYK(t *testing.T) {
	cases := []struct {
		colors   []*color.CMYK
		expected *color.CMYK
	}{
		{
			colors:   []*color.CMYK{},
			expected: &color.CMYK{0, 0, 0, 0},
		},
		{
			colors:   []*color.CMYK{{20, 40, 30, 10}, {10, 50, 20, 20}, {30, 30, 40, 0}},
			expected: &color.CMYK{20, 40, 30, 10},
		},
	}

	for _, c := range cases {
		actual := color.MixCMYK(c.colors...)
		assert.Equal(t, c.expected, actual)
	}
}

func ExampleMixCMYK() {
	var (
		cyan    = color.NewCMYK(100, 0, 0, 40)
		magenta = color.NewCMYK(0, 100, 0, 20)
		yellow  = color.NewCMYK(0, 0, 100, 60)
	)

	grey := color.MixCMYK(cyan, magenta, yellow)

	fmt.Printf("cmyk(%d%%, %d%%, %d%%, %d%%)", grey.C, grey.M, grey.Y, grey.K)
	// Output: cmyk(33%, 33%, 33%, 40%)
}

func TestMixHSL(t *testing.T) {
	cases := []struct {
		colors   []*color.HSL
		expected *color.HSL
	}{
		{
			[]*color.HSL{},
			&color.HSL{0, 0, 0},
		},
		{
			[]*color.HSL{{0, 100, 50}, {120, 100, 50}, {240, 100, 50}},
			&color.HSL{117, 100, 50},
		},
		{
			[]*color.HSL{{221, 30, 43}, {28, 100, 50}},
			&color.HSL{305, 65, 47},
		},
		{
			[]*color.HSL{{356, 75, 43}, {47, 92, 51}, {140, 61, 55}},
			&color.HSL{55, 76, 50},
		},
		{
			[]*color.HSL{{300, 10, 47}, {13, 82, 43}, {203, 98, 44}, {105, 69, 14}},
			&color.HSL{347, 65, 37},
		},
		{
			[]*color.HSL{{228, 58, 68}, {45, 68, 39}, {253, 85, 48}, {155, 83, 70}, {99, 93, 58}},
			&color.HSL{162, 77, 57},
		},
	}

	for _, c := range cases {
		actual := color.MixHSL(c.colors...)
		assert.Equal(t, c.expected, actual)
	}
}

func TestMixRGB(t *testing.T) {
	cases := []struct {
		colors   []*color.RGB
		expected *color.RGB
	}{
		{
			[]*color.RGB{},
			&color.RGB{0, 0, 0},
		},
		{
			[]*color.RGB{{255, 0, 0}, {0, 255, 0}, {0, 0, 255}},
			&color.RGB{85, 85, 85},
		},
		{
			[]*color.RGB{{255, 255, 0}, {0, 255, 255}},
			&color.RGB{128, 255, 128},
		},
		{
			[]*color.RGB{{70, 167, 192}, {0, 255, 0}, {255, 0, 0}},
			&color.RGB{108, 141, 64},
		},
	}

	for _, c := range cases {
		actual := color.MixRGB(c.colors...)
		assert.Equal(t, c.expected, actual)
	}
}

func ExampleMixRGB() {
	var (
		red   = color.NewRGB(255, 0, 0)
		green = color.NewRGB(0, 255, 0)
		blue  = color.NewRGB(0, 0, 255)
	)

	grey := color.MixRGB(red, green, blue)

	fmt.Printf("rgb(%d, %d, %d)\n", grey.R, grey.G, green.B)

}
