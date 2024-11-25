package color_test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cases := []struct {
		r        uint8
		g        uint8
		b        uint8
		expected *color.Color
	}{
		{r: 0, g: 0, b: 0, expected: &color.Color{0, 0, 0}},
		{r: 1, g: 1, b: 1, expected: &color.Color{1, 1, 1}},
		{r: 100, g: 101, b: 102, expected: &color.Color{100, 101, 102}},
		{r: 255, g: 255, b: 255, expected: &color.Color{255, 255, 255}},
	}

	for _, c := range cases {
		actual := color.New(c.r, c.g, c.b)
		assert.EqualExportedValues(t, c.expected, actual)
	}

	// Random tests.
	for range 1000 {
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)

		expected := &color.Color{uint8(r), uint8(g), uint8(b)}
		actual := color.New(uint8(r), uint8(g), uint8(b))

		assert.EqualExportedValues(t, expected, actual)
	}
}

func ExampleNew() {
	green := color.New(0, 255, 0)

	fmt.Println(green.Hex()) // Output: #00ff00
}

func TestNewFromHex(t *testing.T) {
	cases := []struct {
		hex      string
		expected *color.Color
		err      bool
	}{
		{hex: "000", expected: &color.Color{0, 0, 0}, err: false},
		{hex: "#000", expected: &color.Color{0, 0, 0}, err: false},
		{hex: "000000", expected: &color.Color{0, 0, 0}, err: false},
		{hex: "#000000", expected: &color.Color{0, 0, 0}, err: false},
		{hex: "fff", expected: &color.Color{255, 255, 255}, err: false},
		{hex: "#fff", expected: &color.Color{255, 255, 255}, err: false},
		{hex: "ffffff", expected: &color.Color{255, 255, 255}, err: false},
		{hex: "#ffffff", expected: &color.Color{255, 255, 255}, err: false},

		{hex: "ggggggg", err: true},
		{hex: "a", err: true},
		{hex: "not a valid hexadecimal string", err: true},
		{hex: "#######", err: true},
	}

	for _, c := range cases {
		actual, err := color.NewFromHex(c.hex)

		if c.err {
			assert.Nil(t, actual)
			assert.Error(t, err)
		} else {
			assert.EqualExportedValues(t, c.expected, actual)
			assert.NoError(t, err)
		}
	}

	// Random tests.
	for range 1000 {
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)

		hex := fmt.Sprintf("%02x%02x%02x", r, g, b)

		expected := &color.Color{uint8(r), uint8(g), uint8(b)}

		actual, err := color.NewFromHex(hex)
		assert.EqualExportedValues(t, expected, actual)
		assert.Nil(t, err)

		hex = fmt.Sprintf("#%02x%02x%02x", r, g, b)
		t.Log(hex)
		actual, err = color.NewFromHex(hex)
		assert.EqualExportedValues(t, expected, actual)
		assert.Nil(t, err)
	}
}

func ExampleNewFromHex() {
	c, err := color.NewFromHex("#fe8019")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("rgb(%d, %d, %d)\n", c.R, c.G, c.B) // Output: rgb(254, 128, 25)
}

func TestNewFromCMYK(t *testing.T) {
	cases := []struct {
		cyan, magenta, yellow, key int
		expected                   *color.Color
		err                        bool
	}{
		{0, 0, 0, 100, &color.Color{0, 0, 0}, false},
		{0, 0, 0, 0, &color.Color{255, 255, 255}, false},
		{0, 100, 100, 0, &color.Color{255, 0, 0}, false},
		{100, 0, 100, 0, &color.Color{0, 255, 0}, false},
		{100, 100, 0, 0, &color.Color{0, 0, 255}, false},
		{0, 0, 100, 0, &color.Color{255, 255, 0}, false},
		{0, 100, 0, 0, &color.Color{255, 0, 255}, false},
		{100, 0, 0, 0, &color.Color{0, 255, 255}, false},
	}

	for _, c := range cases {
		actual, err := color.NewFromCMYK(c.cyan, c.magenta, c.yellow, c.key)

		if c.err {
			assert.Nil(t, c.expected)
			assert.Error(t, err)
		} else {
			assert.EqualExportedValues(t, c.expected, actual)
			assert.NoError(t, err)
		}
	}
}

func ExampleNewFromCMYK() {
	magenta, err := color.NewFromCMYK(0, 100, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(magenta.Hex()) // Output: #ff00ff
}

func TestNewFromHSL(t *testing.T) {
	cases := []struct {
		hue, saturation, lightness int
		expected                   *color.Color
		err                        bool
	}{
		{0, 100, 50, &color.Color{255, 0, 0}, false},
		{120, 100, 50, &color.Color{0, 255, 0}, false},
		{240, 100, 50, &color.Color{0, 0, 255}, false},
		{60, 100, 50, &color.Color{255, 255, 0}, false},
		{180, 100, 50, &color.Color{0, 255, 255}, false},
		{300, 100, 50, &color.Color{255, 0, 255}, false},
		{0, 0, 50, &color.Color{128, 128, 128}, false},
		{0, 0, 0, &color.Color{0, 0, 0}, false},
		{0, 0, 100, &color.Color{255, 255, 255}, false},
		{8000, 0, 0, nil, true},
	}

	for _, c := range cases {
		actual, err := color.NewFromHSL(c.hue, c.saturation, c.lightness)

		if c.err {
			assert.Nil(t, c.expected)
			assert.Error(t, err)
		} else {
			assert.EqualExportedValues(t, c.expected, actual)
			assert.NoError(t, err)
		}
	}
}

func ExampleNewFromHSL() {
	c, err := color.NewFromHSL(83, 34, 63)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Hex()) // Output: #a8c181
}

func TestColor_Hex(t *testing.T) {
	cases := []struct {
		color    *color.Color
		expected string
	}{
		{color: &color.Color{0, 0, 0}, expected: "#000000"},
		{color: &color.Color{0, 0, 1}, expected: "#000001"},
		{color: &color.Color{1, 2, 3}, expected: "#010203"},
		{color: &color.Color{255, 255, 255}, expected: "#ffffff"},
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
		actual := color.Color{uint8(r), uint8(g), uint8(b)}.Hex()
		assert.Equal(t, expected, actual)
	}
}

func ExampleColor_Hex() {
	black := color.New(0, 0, 0)
	fmt.Println(black.Hex()) // Output: #000000
}

func TestColor_CMYK(t *testing.T) {
	cases := []struct {
		color    *color.Color
		expected *color.CMYK
	}{
		{color: &color.Color{0, 0, 0}, expected: &color.CMYK{0, 0, 0, 100}},
		{color: &color.Color{255, 255, 255}, expected: &color.CMYK{0, 0, 0, 0}},
		{color: &color.Color{255, 0, 0}, expected: &color.CMYK{0, 100, 100, 0}},
		{color: &color.Color{0, 255, 0}, expected: &color.CMYK{100, 0, 100, 0}},
		{color: &color.Color{0, 0, 255}, expected: &color.CMYK{100, 100, 0, 0}},
		{color: &color.Color{255, 255, 0}, expected: &color.CMYK{0, 0, 100, 0}},
		{color: &color.Color{255, 0, 255}, expected: &color.CMYK{0, 100, 0, 0}},
		{color: &color.Color{0, 255, 255}, expected: &color.CMYK{100, 0, 0, 0}},
	}

	for _, c := range cases {
		actual := c.color.CMYK()
		assert.EqualExportedValues(t, c.expected, actual)
	}
}

func ExampleColor_CMYK() {
	yellow := color.New(255, 255, 0)
	cmyk := yellow.CMYK()
	fmt.Printf("cmyk(%d%%, %d%%, %d%%, %d%%)", cmyk.C, cmyk.M, cmyk.Y, cmyk.K) // Output: cmyk(0%, 0%, 100%, 0%)
}

func TestColor_HSL(t *testing.T) {
	cases := []struct {
		color    *color.Color
		expected *color.HSL
	}{
		{color: &color.Color{255, 0, 0}, expected: &color.HSL{0, 100, 50}},
		{color: &color.Color{0, 255, 0}, expected: &color.HSL{120, 100, 50}},
		{color: &color.Color{0, 0, 255}, expected: &color.HSL{240, 100, 50}},
		{color: &color.Color{255, 255, 0}, expected: &color.HSL{60, 100, 50}},
		{color: &color.Color{0, 255, 255}, expected: &color.HSL{180, 100, 50}},
		{color: &color.Color{255, 0, 255}, expected: &color.HSL{300, 100, 50}},
		{color: &color.Color{128, 128, 128}, expected: &color.HSL{0, 0, 50}},
		{color: &color.Color{0, 0, 0}, expected: &color.HSL{0, 0, 0}},
		{color: &color.Color{255, 255, 255}, expected: &color.HSL{0, 0, 100}},
		{color: &color.Color{184, 201, 221}, expected: &color.HSL{212, 35, 79}},
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
		color := color.Color{uint8(r), uint8(g), uint8(b)}
		hsl := color.HSL()

		assert.True(t, hsl.Hue >= 0 && hsl.Hue < 360)
		assert.True(t, hsl.Saturation >= 0 && hsl.Saturation <= 100)
		assert.True(t, hsl.Lightness >= 0 && hsl.Lightness <= 100)
	}
}

func ExampleColor_HSL() {
	c := color.New(219, 188, 127)
	hsl := c.HSL()
	fmt.Printf("hsl(%d, %d%%, %d%%)\n", hsl.Hue, hsl.Saturation, hsl.Lightness) // Output: hsl(39, 56%, 67%)
}
