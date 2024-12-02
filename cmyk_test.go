package color_test

import (
	"fmt"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestCMYK(t *testing.T) {
	assert.Implements(t, (*interface {
		Hex() string
		HSL() *color.HSL
		HSV() *color.HSV
		RGB() *color.RGB
		XYZ() *color.XYZ

		String() string
	})(nil), color.CMYK{})
}

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

func TestCMYK_Edit(t *testing.T) {
	cases := []struct {
		color    *color.CMYK
		expected *color.CMYK
		editfn   func(*color.CMYK)
	}{
		{
			color:    &color.CMYK{0, 0, 0, 100},
			expected: &color.CMYK{0, 0, 0, 0},
			editfn: func(c *color.CMYK) {
				c.K = 0
			},
		},
		{
			color:    &color.CMYK{1, 2, 4, 8},
			expected: &color.CMYK{2, 4, 8, 16},
			editfn: func(c *color.CMYK) {
				c.C *= 2
				c.M *= 2
				c.Y *= 2
				c.K *= 2
			},
		},
	}

	for _, c := range cases {
		actual := c.color.Edit(c.editfn)
		assert.Equal(t, c.expected, actual)
		assert.Equal(t, c.expected, c.color)
	}
}

func ExampleCMYK_Edit() {
	orange := color.NewCMYK(0, 47, 90, 0) // rgb(255, 134, 26)

	// Make it darker and print RGB.
	fmt.Println(orange.Edit(func(c *color.CMYK) {
		c.K = 50
	}).RGB()) // Output: rgb(128, 68, 13)
}
