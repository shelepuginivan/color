package color_test

import (
	"fmt"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestCMYK_RGB(t *testing.T) {
	cases := []struct {
		color    *color.CMYK
		expected *color.RGB
	}{
		{&color.CMYK{0, 0, 0, 100}, &color.RGB{0, 0, 0}},
		{&color.CMYK{0, 0, 0, 0}, &color.RGB{255, 255, 255}},
		{&color.CMYK{0, 100, 100, 0}, &color.RGB{255, 0, 0}},
		{&color.CMYK{100, 0, 100, 0}, &color.RGB{0, 255, 0}},
		{&color.CMYK{100, 100, 0, 0}, &color.RGB{0, 0, 255}},
		{&color.CMYK{0, 0, 100, 0}, &color.RGB{255, 255, 0}},
		{&color.CMYK{0, 100, 0, 0}, &color.RGB{255, 0, 255}},
		{&color.CMYK{100, 0, 0, 0}, &color.RGB{0, 255, 255}},
	}

	for _, c := range cases {
		actual := c.color.RGB()
		assert.Equal(t, c.expected, actual)
	}
}

func ExampleCMYK_RGB() {
	yellow := color.NewCMYK(0, 0, 100, 0)
	fmt.Println(yellow.RGB()) // Output: rgb(255, 255, 0)
}
