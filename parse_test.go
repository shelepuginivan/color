package color_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shelepuginivan/color"

	"github.com/stretchr/testify/assert"
)

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

		expected := &color.RGB{uint8(r), uint8(g), uint8(b)}

		actual, err := color.ParseHex(hex)
		assert.EqualExportedValues(t, expected, actual)
		assert.Nil(t, err)

		hex = fmt.Sprintf("#%02x%02x%02x", r, g, b)
		t.Log(hex)
		actual, err = color.ParseHex(hex)
		assert.EqualExportedValues(t, expected, actual)
		assert.Nil(t, err)
	}
}
