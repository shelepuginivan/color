package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestOklab(t *testing.T) {
	assert.Implements(t, (*color.Color)(nil), color.Oklab{})
}

func TestOklab_XYZ(t *testing.T) {
	cases := []struct {
		color    *color.Oklab
		expected *color.XYZ
	}{
		{&color.Oklab{0.51634019, 0.15469500, 0.06289579}, &color.XYZ{0.2065400, 0.1219722, 0.0513695}},
	}

	for _, c := range cases {
		actual := c.color.XYZ()

		assert.InDelta(t, c.expected.X, actual.X, 0.05)
		assert.InDelta(t, c.expected.Y, actual.Y, 0.05)
		assert.InDelta(t, c.expected.Z, actual.Z, 0.05)
	}
}

func TestOklab_Oklch(t *testing.T) {
	cases := []struct {
		color    *color.Oklab
		expected *color.Oklch
	}{
		{&color.Oklab{0.6280, 0.2248, 0.1258}, &color.Oklch{0.6280, 0.2576, 29}},
	}

	for _, c := range cases {
		actual := c.color.Oklch()

		assert.InDelta(t, c.expected.L, actual.L, 0.001)
		assert.InDelta(t, c.expected.C, actual.C, 0.001)
		assert.Equal(t, c.expected.H, actual.H)
	}
}
