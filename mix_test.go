package color_test

import (
	"fmt"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestMixCMYK(t *testing.T) {
	cases := []struct {
		colors   []color.Color
		expected *color.CMYK
	}{
		{
			colors:   []color.Color{},
			expected: &color.CMYK{0, 0, 0, 0},
		},
		{
			colors: []color.Color{
				&color.CMYK{20, 40, 30, 10},
				&color.CMYK{10, 50, 20, 20},
				&color.CMYK{30, 30, 40, 0},
			},
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
		colors   []color.Color
		expected *color.HSL
	}{
		{
			[]color.Color{},
			&color.HSL{0, 0, 0},
		},
		{
			[]color.Color{
				&color.HSL{0, 100, 50},
				&color.HSL{120, 100, 50},
				&color.HSL{240, 100, 50},
			},
			&color.HSL{117, 100, 50},
		},
		{
			[]color.Color{
				&color.HSL{221, 30, 43},
				&color.HSL{28, 100, 50},
			},
			&color.HSL{305, 65, 47},
		},
		{
			[]color.Color{
				&color.HSL{356, 75, 43},
				&color.HSL{47, 92, 51},
				&color.HSL{140, 61, 55},
			},
			&color.HSL{55, 76, 50},
		},
		{
			[]color.Color{
				&color.HSL{300, 10, 47},
				&color.HSL{13, 82, 43},
				&color.HSL{203, 98, 44},
				&color.HSL{105, 69, 14},
			},
			&color.HSL{347, 65, 37},
		},
		{
			[]color.Color{
				&color.HSL{228, 58, 68},
				&color.HSL{45, 68, 39},
				&color.HSL{253, 85, 48},
				&color.HSL{155, 83, 70},
				&color.HSL{99, 93, 58},
			},
			&color.HSL{162, 77, 57},
		},
	}

	for _, c := range cases {
		actual := color.MixHSL(c.colors...)
		assert.Equal(t, c.expected, actual)
	}
}

func TestMixLab(t *testing.T) {
	cases := []struct {
		colors   []color.Color
		expected *color.Lab
	}{
		{
			[]color.Color{},
			&color.Lab{0, 0, 0},
		},
		{
			[]color.Color{
				&color.Lab{39.73, -4.86, -25.62},
				&color.Lab{23.73, 30.23, -56.88},
				&color.Lab{57.83, 82.11, -49.95},
			},
			&color.Lab{40.43, 35.83, -44.15},
		},
		{
			[]color.Color{
				&color.Lab{70.13, 32.87, 74.55},
				&color.Lab{86.30, -8.82, 76.56},
				&color.Lab{81.03, 18.75, 50.84},
			},
			&color.Lab{79.15, 14.29, 67.32},
		},
		{
			[]color.Color{
				&color.Lab{10.39, 22.23, -43.88},
				&color.Lab{54.31, 63.08, -23.74},
				&color.Lab{61.92, 51.14, 68.16},
				&color.Lab{38.45, -0.82, 37.27},
			},
			&color.Lab{41.27, 33.91, 9.45},
		},
	}

	for _, c := range cases {
		actual := color.MixLab(c.colors...)

		assert.InDelta(t, c.expected.L, actual.L, 0.05)
		assert.InDelta(t, c.expected.A, actual.A, 0.05)
		assert.InDelta(t, c.expected.B, actual.B, 0.05)
	}
}

func TestMixRGB(t *testing.T) {
	cases := []struct {
		colors   []color.Color
		expected *color.RGB
	}{
		{
			[]color.Color{},
			&color.RGB{0, 0, 0},
		},
		{
			[]color.Color{
				&color.RGB{255, 0, 0},
				&color.RGB{0, 255, 0},
				&color.RGB{0, 0, 255},
			},
			&color.RGB{85, 85, 85},
		},
		{
			[]color.Color{
				&color.RGB{255, 255, 0},
				&color.RGB{0, 255, 255},
			},
			&color.RGB{128, 255, 128},
		},
		{
			[]color.Color{
				&color.RGB{70, 167, 192},
				&color.RGB{0, 255, 0},
				&color.RGB{255, 0, 0},
			},
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
