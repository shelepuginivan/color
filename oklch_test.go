package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestOklch(t *testing.T) {
	assert.Implements(t, (*color.Color)(nil), color.Oklch{})
}

func TestOklch_Oklab(t *testing.T) {
	cases := []struct {
		color    *color.Oklch
		expected *color.Oklab
	}{
		{&color.Oklch{0.5428, 0.0526, 157}, &color.Oklab{0.5428, -0.0485, 0.0206}},
	}

	for _, c := range cases {
		actual := c.color.Oklab()

		assert.InDelta(t, c.expected.L, actual.L, 0.005)
		assert.InDelta(t, c.expected.A, actual.A, 0.005)
		assert.InDelta(t, c.expected.B, actual.B, 0.005)
	}
}
