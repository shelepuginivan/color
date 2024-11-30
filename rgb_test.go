package color_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestRGB(t *testing.T) {
	assert.Implements(t, (*interface {
		CMYK() *color.CMYK
		Hex() string
		HSL() *color.HSL
		HSV() *color.HSV

		String() string
	})(nil), color.RGB{})
}

func TestNewRGB(t *testing.T) {
	cases := []struct {
		r        uint8
		g        uint8
		b        uint8
		expected *color.RGB
	}{
		{r: 0, g: 0, b: 0, expected: &color.RGB{0, 0, 0}},
		{r: 1, g: 1, b: 1, expected: &color.RGB{1, 1, 1}},
		{r: 100, g: 101, b: 102, expected: &color.RGB{100, 101, 102}},
		{r: 255, g: 255, b: 255, expected: &color.RGB{255, 255, 255}},
	}

	for _, c := range cases {
		actual := color.NewRGB(c.r, c.g, c.b)
		assert.EqualExportedValues(t, c.expected, actual)
	}

	// Random tests.
	for range 1000 {
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)

		expected := &color.RGB{uint8(r), uint8(g), uint8(b)}
		actual := color.NewRGB(uint8(r), uint8(g), uint8(b))

		assert.EqualExportedValues(t, expected, actual)
	}
}

func ExampleNewRGB() {
	green := color.NewRGB(0, 255, 0)

	fmt.Println(green.Hex()) // Output: #00ff00
}

func TestRGB_Hex(t *testing.T) {
	cases := []struct {
		color    *color.RGB
		expected string
	}{
		{color: &color.RGB{0, 0, 0}, expected: "#000000"},
		{color: &color.RGB{0, 0, 1}, expected: "#000001"},
		{color: &color.RGB{1, 2, 3}, expected: "#010203"},
		{color: &color.RGB{255, 255, 255}, expected: "#ffffff"},
	}

	for _, c := range cases {
		actual := c.color.Hex()
		assert.Equal(t, c.expected, actual)
	}

	// Random tests.
	for range 1000 {
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)
		expected := fmt.Sprintf("#%02x%02x%02x", r, g, b)
		actual := color.RGB{uint8(r), uint8(g), uint8(b)}.Hex()
		assert.Equal(t, expected, actual)
	}
}

func ExampleRGB_Hex() {
	black := color.NewRGB(0, 0, 0)
	fmt.Println(black.Hex()) // Output: #000000
}

func TestRGB_CMYK(t *testing.T) {
	cases := []struct {
		color    *color.RGB
		expected *color.CMYK
	}{
		{color: &color.RGB{0, 0, 0}, expected: &color.CMYK{0, 0, 0, 100}},
		{color: &color.RGB{255, 255, 255}, expected: &color.CMYK{0, 0, 0, 0}},
		{color: &color.RGB{255, 0, 0}, expected: &color.CMYK{0, 100, 100, 0}},
		{color: &color.RGB{0, 255, 0}, expected: &color.CMYK{100, 0, 100, 0}},
		{color: &color.RGB{0, 0, 255}, expected: &color.CMYK{100, 100, 0, 0}},
		{color: &color.RGB{255, 255, 0}, expected: &color.CMYK{0, 0, 100, 0}},
		{color: &color.RGB{255, 0, 255}, expected: &color.CMYK{0, 100, 0, 0}},
		{color: &color.RGB{0, 255, 255}, expected: &color.CMYK{100, 0, 0, 0}},
	}

	for _, c := range cases {
		actual := c.color.CMYK()
		assert.EqualExportedValues(t, c.expected, actual)
	}
}

func ExampleRGB_CMYK() {
	yellow := color.NewRGB(255, 255, 0)
	cmyk := yellow.CMYK()
	fmt.Printf("cmyk(%d%%, %d%%, %d%%, %d%%)", cmyk.C, cmyk.M, cmyk.Y, cmyk.K) // Output: cmyk(0%, 0%, 100%, 0%)
}

func TestColor_HSL(t *testing.T) {
	cases := []struct {
		color    *color.RGB
		expected *color.HSL
	}{
		{color: &color.RGB{255, 0, 0}, expected: &color.HSL{0, 100, 50}},
		{color: &color.RGB{0, 255, 0}, expected: &color.HSL{120, 100, 50}},
		{color: &color.RGB{0, 0, 255}, expected: &color.HSL{240, 100, 50}},
		{color: &color.RGB{255, 255, 0}, expected: &color.HSL{60, 100, 50}},
		{color: &color.RGB{0, 255, 255}, expected: &color.HSL{180, 100, 50}},
		{color: &color.RGB{255, 0, 255}, expected: &color.HSL{300, 100, 50}},
		{color: &color.RGB{128, 128, 128}, expected: &color.HSL{0, 0, 50}},
		{color: &color.RGB{0, 0, 0}, expected: &color.HSL{0, 0, 0}},
		{color: &color.RGB{255, 255, 255}, expected: &color.HSL{0, 0, 100}},
		{color: &color.RGB{184, 201, 221}, expected: &color.HSL{212, 35, 79}},
	}

	for _, c := range cases {
		actual := c.color.HSL()
		assert.Equal(t, c.expected, actual)
	}

	// Random tests.
	for range 1000 {
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)
		color := color.RGB{uint8(r), uint8(g), uint8(b)}
		hsl := color.HSL()

		assert.True(t, hsl.H >= 0 && hsl.H < 360)
		assert.True(t, hsl.S >= 0 && hsl.S <= 100)
		assert.True(t, hsl.L >= 0 && hsl.L <= 100)
	}
}

func ExampleRGB_HSL() {
	c := color.NewRGB(219, 188, 127)
	fmt.Println(c.HSL()) // Output: hsl(40, 56%, 68%)
}

func TestRGB_HSV(t *testing.T) {
	cases := []struct {
		color    *color.RGB
		expected *color.HSV
	}{
		{&color.RGB{255, 0, 0}, &color.HSV{0, 100, 100}},
		{&color.RGB{0, 255, 0}, &color.HSV{120, 100, 100}},
		{&color.RGB{0, 0, 255}, &color.HSV{240, 100, 100}},
		{&color.RGB{255, 255, 0}, &color.HSV{60, 100, 100}},
		{&color.RGB{0, 255, 255}, &color.HSV{180, 100, 100}},
		{&color.RGB{255, 0, 255}, &color.HSV{300, 100, 100}},
		{&color.RGB{128, 128, 128}, &color.HSV{0, 0, 50}},
		{&color.RGB{0, 0, 0}, &color.HSV{0, 0, 0}},
		{&color.RGB{255, 255, 255}, &color.HSV{0, 0, 100}},
		{&color.RGB{0, 0, 128}, &color.HSV{240, 100, 50}},
		{&color.RGB{0, 128, 128}, &color.HSV{180, 100, 50}},
		{&color.RGB{128, 0, 0}, &color.HSV{0, 100, 50}},
	}

	for _, c := range cases {
		actual := c.color.HSV()
		assert.Equal(t, c.expected, actual)
	}

	// Random tests.
	for range 1000 {
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)
		color := color.RGB{uint8(r), uint8(g), uint8(b)}
		hsv := color.HSV()

		assert.True(t, hsv.H >= 0 && hsv.H < 360)
		assert.True(t, hsv.S >= 0 && hsv.S <= 100)
		assert.True(t, hsv.V >= 0 && hsv.V <= 100)
	}
}

func TestRGB_Edit(t *testing.T) {
	cases := []struct {
		color    *color.RGB
		expected *color.RGB
		editfn   func(*color.RGB)
	}{
		{
			color:    &color.RGB{200, 200, 200},
			expected: &color.RGB{100, 100, 200},
			editfn: func(c *color.RGB) {
				c.R /= 2
				c.G /= 2
			},
		},
		{
			color:    &color.RGB{255, 0, 0},
			expected: &color.RGB{0, 255, 0},
			editfn: func(c *color.RGB) {
				c.R, c.G = c.G, c.R
			},
		},
	}

	for _, c := range cases {
		actual := c.color.Edit(c.editfn)
		assert.Equal(t, c.expected, actual)
		assert.Equal(t, c.expected, c.color)
	}
}

func ExampleRGB_Edit() {
	silver := color.NewRGB(191, 191, 191)
	fmt.Println(silver.Edit(func(c *color.RGB) {
		c.R += 10
		c.G += 20
		c.B += 30
	})) // Output: rgb(201, 211, 221)
}
