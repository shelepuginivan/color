package color_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestParseNamed(t *testing.T) {
	cases := []struct {
		input    string
		expected *color.RGB
		isErr    bool
	}{
		{"maroon", &color.RGB{128, 0, 0}, false},
		{"red", &color.RGB{255, 0, 0}, false},
		{"white", &color.RGB{255, 255, 255}, false},
		{"black", &color.RGB{0, 0, 0}, false},
		{"ReD", &color.RGB{255, 0, 0}, false},
		{"notacolor", nil, true},
		{"", nil, true},
	}

	for _, c := range cases {
		actual, err := color.ParseNamed(c.input)

		if c.isErr {
			assert.Nil(t, c.expected)
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.expected, actual)
		}
	}
}

func TestParseHex(t *testing.T) {
	cases := []struct {
		hex      string
		expected *color.RGB
		err      bool
	}{
		{hex: "000", expected: &color.RGB{0, 0, 0}, err: false},
		{hex: "#000", expected: &color.RGB{0, 0, 0}, err: false},
		{hex: "000000", expected: &color.RGB{0, 0, 0}, err: false},
		{hex: "#000000", expected: &color.RGB{0, 0, 0}, err: false},
		{hex: "fff", expected: &color.RGB{255, 255, 255}, err: false},
		{hex: "#fff", expected: &color.RGB{255, 255, 255}, err: false},
		{hex: "ffffff", expected: &color.RGB{255, 255, 255}, err: false},
		{hex: "#ffffff", expected: &color.RGB{255, 255, 255}, err: false},

		{hex: "ggggggg", err: true},
		{hex: "a", err: true},
		{hex: "ab", err: true},
		{hex: "ffv", err: true},
		{hex: "not a valid hexadecimal string", err: true},
		{hex: "#######", err: true},
	}

	for _, c := range cases {
		actual, err := color.ParseHex(c.hex)

		if c.err {
			assert.Nil(t, actual)
			assert.Error(t, err)
		} else {
			assert.Equal(t, c.expected, actual)
			assert.NoError(t, err)
		}
	}

	// Random tests.
	for range 1000 {
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)

		hex := fmt.Sprintf("%02x%02x%02x", r, g, b)

		expected := &color.RGB{uint8(r), uint8(g), uint8(b)}

		actual, err := color.ParseHex(hex)
		assert.EqualExportedValues(t, expected, actual)
		assert.Nil(t, err)

		hex = fmt.Sprintf("#%02x%02x%02x", r, g, b)

		actual, err = color.ParseHex(hex)
		assert.EqualExportedValues(t, expected, actual)
		assert.Nil(t, err)
	}
}

func TestParseFunc(t *testing.T) {
	cases := []struct {
		fnstring string
		expected color.Color
		err      bool
	}{
		{"", nil, true},
		{"erkeropgjerpgjerg", nil, true},

		{"cmyk(0%, 18%, 39%, 5%)", &color.CMYK{0, 18, 39, 5}, false},
		{"cmyk(61% 6% 0% 35%)", &color.CMYK{61, 6, 0, 35}, false},
		{"cmyk(0.5, 0.5, 0.5, 0.5)", &color.CMYK{50, 50, 50, 50}, false},
		{"cmyk(.65 .34 .32 .11)", &color.CMYK{65, 34, 32, 11}, false},
		{"cmyk(.03 .33 .24 none)", &color.CMYK{3, 33, 24, 0}, false},
		{"cmyk(not a valid arg)", nil, true},
		{"cmyk(65% 19% 78%)", nil, true},
		{"cmyk(34% 21% 98% 67% 45%)", nil, true},

		{"hsl(112, 92%, 43%)", &color.HSL{112, 92, 43}, false},
		{"hsl(221 87% 53%)", &color.HSL{221, 87, 53}, false},
		{"hsl(288deg .98, 82%)", &color.HSL{288, 98, 82}, false},
		{"hsl(0.3turn 60% 45%)", &color.HSL{108, 60, 45}, false},
		{"hsl(0.22689rad 97% 59%)", &color.HSL{13, 97, 59}, false},
		{"hsl(a, b, c)", nil, true},
		{"hsl(67 .83)", nil, true},
		{"hsl(310 79% 66% .17)", nil, true},

		{"hsv(343, 80%, 90%)", &color.HSV{343, 80, 90}, false},
		{"hsv(209 .72 .83)", &color.HSV{209, 72, 83}, false},
		{"hsv(132deg 92% 77%)", &color.HSV{132, 92, 77}, false},
		{"hsv(0.767turn, 99%, 71%)", &color.HSV{276, 99, 71}, false},
		{"hsv(1.9199rad 9% 49%)", &color.HSV{110, 9, 49}, false},
		{"hsv(a, b, c)", nil, true},
		{"hsv(67 .83)", nil, true},
		{"hsv(310 79% 66% .17)", nil, true},

		{"lab(53.27, 80.109, 67.220)", &color.Lab{53.27, 80.109, 67.220}, false},
		{"lab(87.73 -86.1846 83.181)", &color.Lab{87.73, -86.1846, 83.181}, false},
		{"lab(100.0, 0.526%, -1.04%)", &color.Lab{100.0, 0.00526, -0.0104}, false},
		{"lab(string string string)", nil, true},
		{"lab(100.0, 86%)", nil, true},
		{"lab(1 1 1 1)", nil, true},

		{"lch(8.991815706465342, 3.7396951758251333, 82)", &color.Lch{8.991815706465342, 3.7396951758251333, 82}, false},
		{"lch(3.234 1.75672 none)", &color.Lch{3.234, 1.75672, 0}, false},
		{"lch(again invalid args)", nil, true},
		{"lch(3.141592654 2.718281828)", nil, true},
		{"lch(0 0 0 0)", nil, true},

		{"oklab(0.51634019, 0.15469500, 0.06289579)", &color.Oklab{0.51634019, 0.15469500, 0.06289579}, false},
		{"oklab(0.20654008 0.12197225 0.05136952)", &color.Oklab{0.20654008, 0.12197225, 0.05136952}, false},
		{"oklab(0.3842619, none, 18.88683%)", &color.Oklab{0.3842619, 0, 0.1888683}, false},
		{"oklab(yet another case)", nil, true},
		{"oklab(1 2 3 4)", nil, true},

		{"oklch(8.991815706465342, 3.7396951758251333, 82)", &color.Oklch{8.991815706465342, 3.7396951758251333, 82}, false},
		{"oklch(3.234 1.75672 0.1turn)", &color.Oklch{3.234, 1.75672, 36}, false},
		{"oklch(again invalid args)", nil, true},
		{"oklch(3.141592654 2.718281828)", nil, true},
		{"oklch(0 0 0 0)", nil, true},

		{"rgb(255 255 255)", &color.RGB{255, 255, 255}, false},
		{"rgb(none none none)", &color.RGB{0, 0, 0}, false},
		{"rgb(r g b)", nil, true},
		{"rgb(30 40)", nil, true},
		{"rgb(11 22 33 44)", nil, true},

		{"xyz(1.0985, 1.0000, 0.3558)", color.A, false},
		{"xyz(0.9807, 1.0000, 1.1822)", color.C, false},
		{"xyz(0.9505 1 1.0888)", color.D65, false},
		{"xyz(33% 99% 101%)", &color.XYZ{0.33, 0.99, 1.01}, false},
		{"xyz(x y z)", nil, true},
		{"xyz(30% 20%)", nil, true},
		{"xyz(.11 .12 .13 .14)", nil, true},
	}

	for _, c := range cases {
		actual, err := color.ParseFunc(c.fnstring)

		if c.err {
			assert.Nil(t, actual)
			assert.Error(t, err)
		} else {
			assert.Equal(t, c.expected, actual)
			assert.NoError(t, err)
		}
	}
}
