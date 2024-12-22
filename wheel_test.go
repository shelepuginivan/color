package color_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/shelepuginivan/color/internal/degrees"
	"github.com/stretchr/testify/assert"
)

func TestComplementary(t *testing.T) {
	for range 1000 {
		var (
			h = rand.IntN(360)
			s = rand.IntN(101)
			l = rand.IntN(101)

			c = &color.HSL{h, s, l}
		)

		actual := color.Complementary(c).HSL()

		assert.Equal(t, degrees.Normalize(h+180), actual.H)
		assert.Equal(t, s, actual.S)
		assert.Equal(t, l, actual.L)
	}
}

func ExampleComplementary() {
	blue := color.Must(color.ParseHex("#0000ff"))

	// Yellow is the complementary color for blue.
	yellow := color.Complementary(blue)

	fmt.Println(yellow.Hex()) // Output: #ffff00
}

func TestSplitComplementary(t *testing.T) {
	for range 1000 {
		var (
			h = rand.IntN(360)
			s = rand.IntN(101)
			l = rand.IntN(101)

			c = &color.HSL{h, s, l}
		)

		var (
			c1, c2 = color.SplitComplementary(c)
			a      = c1.HSL()
			b      = c2.HSL()
		)

		assert.Equal(t, degrees.Normalize(h+150), a.H)
		assert.Equal(t, s, a.S)
		assert.Equal(t, l, a.L)

		assert.Equal(t, degrees.Normalize(h+210), b.H)
		assert.Equal(t, s, b.S)
		assert.Equal(t, l, b.L)
	}
}

func ExampleSplitComplementary() {
	cyan := color.Must(color.ParseHex("#00ffff"))

	// Crimson and orange are split complementary colors of cyan.
	crimson, orange := color.SplitComplementary(cyan)

	fmt.Println(crimson.Hex())
	fmt.Println(orange.Hex())

	// Output:
	// #ff0080
	// #ff8000
}

func TestTriadic(t *testing.T) {
	for range 1000 {
		var (
			h = rand.IntN(360)
			s = rand.IntN(101)
			l = rand.IntN(101)

			c = &color.HSL{h, s, l}
		)

		var (
			c1, c2 = color.Triadic(c)
			a      = c1.HSL()
			b      = c2.HSL()
		)

		assert.Equal(t, degrees.Normalize(h+120), a.H)
		assert.Equal(t, s, a.S)
		assert.Equal(t, l, a.L)

		assert.Equal(t, degrees.Normalize(h+240), b.H)
		assert.Equal(t, s, b.S)
		assert.Equal(t, l, b.L)
	}
}

func ExampleTriadic() {
	red := color.Must(color.ParseHex("#ff0000"))

	// Green and blue are the colors in the triadic colorscheme with red.
	green, blue := color.Triadic(red)

	fmt.Println(green.Hex())
	fmt.Println(blue.Hex())

	// Output:
	// #00ff00
	// #0000ff
}

func TestTetradic(t *testing.T) {
	for range 1000 {
		var (
			h = rand.IntN(360)
			s = rand.IntN(101)
			l = rand.IntN(101)

			c = &color.HSL{h, s, l}
		)

		var (
			c1, c2, c3 = color.Tetradic(c)
			a          = c1.HSL()
			b          = c2.HSL()
			d          = c3.HSL()
		)

		assert.Equal(t, degrees.Normalize(h+90), a.H)
		assert.Equal(t, s, a.S)
		assert.Equal(t, l, a.L)

		assert.Equal(t, degrees.Normalize(h+180), b.H)
		assert.Equal(t, s, b.S)
		assert.Equal(t, l, b.L)

		assert.Equal(t, degrees.Normalize(h+270), d.H)
		assert.Equal(t, s, d.S)
		assert.Equal(t, l, d.L)
	}
}

func ExampleTetradic() {
	magenta := color.Must(color.ParseHex("ff00ff"))

	// Orange, green and blue are the colors in the tetradic colorscheme with magenta.
	orange, green, blue := color.Tetradic(magenta)

	fmt.Println(orange.Hex())
	fmt.Println(green.Hex())
	fmt.Println(blue.Hex())

	// Output:
	// #ff8000
	// #00ff00
	// #007fff
}

func TestAnalogous(t *testing.T) {
	for range 1000 {
		var (
			h = rand.IntN(360)
			s = rand.IntN(101)
			l = rand.IntN(101)

			c = &color.HSL{h, s, l}
		)

		var (
			c1, c2 = color.Analogous(c)
			a      = c1.HSL()
			b      = c2.HSL()
		)

		assert.Equal(t, degrees.Normalize(h-30), a.H)
		assert.Equal(t, s, a.S)
		assert.Equal(t, l, a.L)

		assert.Equal(t, degrees.Normalize(h+30), b.H)
		assert.Equal(t, s, b.S)
		assert.Equal(t, l, b.L)
	}
}

func ExampleAnalogous() {
	orange := color.Must(color.ParseHex("#ff8000"))

	// Red and yellow are adjacent to orange.
	red, yellow := color.Analogous(orange)

	fmt.Println(red.Hex())
	fmt.Println(yellow.Hex())

	// Output:
	// #ff0000
	// #ffff00
}
