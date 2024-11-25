package color_test

import (
	"fmt"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestMix(t *testing.T) {
	cases := []struct {
		colors   []*color.Color
		expected *color.Color
	}{
		{
			[]*color.Color{},
			&color.Color{0, 0, 0},
		},
		{
			[]*color.Color{{255, 0, 0}, {0, 255, 0}, {0, 0, 255}},
			&color.Color{85, 85, 85},
		},
		{
			[]*color.Color{{255, 255, 0}, {0, 255, 255}},
			&color.Color{128, 255, 128},
		},
		{
			[]*color.Color{{70, 167, 192}, {0, 255, 0}, {255, 0, 0}},
			&color.Color{108, 141, 64},
		},
	}

	for _, c := range cases {
		actual := color.Mix(c.colors...)
		assert.Equal(t, c.expected, actual)
	}
}

func ExampleMix() {
	red := color.New(255, 0, 0)
	green := color.New(0, 255, 0)
	blue := color.New(0, 0, 255)

	mix := color.Mix(red, green, blue)
	fmt.Println(mix.Hex()) // Output: #555555
}
