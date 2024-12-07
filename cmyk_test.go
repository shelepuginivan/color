package color_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestCMYK(t *testing.T) {
	assert.Implements(t, (*color.Color)(nil), color.CMYK{})
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

func TestCMYK_Hex(t *testing.T) {
	cases := []struct {
		color    *color.CMYK
		expected string
	}{
		{&color.CMYK{0, 0, 0, 0}, "#ffffff"},
		{&color.CMYK{0, 0, 0, 100}, "#000000"},
		{&color.CMYK{0, 100, 100, 0}, "#ff0000"},
		{&color.CMYK{100, 0, 100, 0}, "#00ff00"},
		{&color.CMYK{100, 100, 0, 0}, "#0000ff"},
		{&color.CMYK{100, 100, 100, 0}, "#000000"},
		{&color.CMYK{25, 0, 76, 20}, "#99cc31"},
		{&color.CMYK{40, 68, 35, 26}, "#713c7b"},
		{&color.CMYK{57, 33, 34, 12}, "#609694"},
		{&color.CMYK{22, 23, 23, 0}, "#c7c4c4"},
	}

	for _, c := range cases {
		actual := c.color.Hex()
		assert.Equal(t, c.expected, actual)
	}
}

func ExampleCMYK_Hex() {
	coral := color.NewCMYK(0, 50, 70, 0)
	fmt.Println(coral.Hex()) // Output: #ff804d
}

func TestCMYK_HSL(t *testing.T) {
	cases := []struct {
		color    *color.CMYK
		expected *color.HSL
	}{
		{&color.CMYK{0, 0, 0, 100}, &color.HSL{0, 0, 0}},
		{&color.CMYK{0, 0, 0, 0}, &color.HSL{0, 0, 100}},
		{&color.CMYK{0, 100, 100, 0}, &color.HSL{0, 100, 50}},
		{&color.CMYK{100, 0, 100, 0}, &color.HSL{120, 100, 50}},
		{&color.CMYK{100, 100, 0, 0}, &color.HSL{240, 100, 50}},
		{&color.CMYK{0, 0, 100, 0}, &color.HSL{60, 100, 50}},
		{&color.CMYK{0, 100, 0, 0}, &color.HSL{300, 100, 50}},
		{&color.CMYK{100, 0, 0, 0}, &color.HSL{180, 100, 50}},
	}

	for _, c := range cases {
		actual := c.color.HSL()
		assert.Equal(t, c.expected, actual)
	}
}

func ExampleCMYK_HSL() {
	skyblue := color.NewCMYK(43, 12, 0, 8)
	fmt.Println(skyblue.HSL()) // Output: hsl(197, 72%, 72%)
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

func TestCMYK_String(t *testing.T) {
	for range 1000 {
		var (
			c = rand.Intn(101)
			m = rand.Intn(101)
			y = rand.Intn(101)
			k = rand.Intn(101)
		)

		var (
			expected = fmt.Sprintf("cmyk(%d%%, %d%%, %d%%, %d%%)", c, m, y, k)
			actual   = color.NewCMYK(c, m, y, k).String()
		)

		assert.Equal(t, expected, actual)
	}
}
