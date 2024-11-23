package color_test

import (
	"fmt"
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

func TestNewFromHSL(t *testing.T) {
	cases := []struct {
		hue, saturation, lightness int
		expected                   *color.Color
	}{
		{0, 100, 50, &color.Color{255, 0, 0}},
		{120, 100, 50, &color.Color{0, 255, 0}},
		{240, 100, 50, &color.Color{0, 0, 255}},
		{60, 100, 50, &color.Color{255, 255, 0}},
		{180, 100, 50, &color.Color{0, 255, 255}},
		{300, 100, 50, &color.Color{255, 0, 255}},
		{0, 0, 50, &color.Color{128, 128, 128}},
		{0, 0, 0, &color.Color{0, 0, 0}},
		{0, 0, 100, &color.Color{255, 255, 255}},
	}

	for _, c := range cases {
		actual := color.NewFromHSL(c.hue, c.saturation, c.lightness)
		assert.Equal(t, c.expected, actual)
	}
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
