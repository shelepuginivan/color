package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestLch(t *testing.T) {
	assert.Implements(t, (*color.Color)(nil), color.Lch{})
}

func TestLch_Lab(t *testing.T) {
	cases := []struct {
		color    *color.Lch
		expected *color.Lab
	}{
		{
			&color.Lch{8.991815706465342, 3.7396951758251333, 82},
			&color.Lab{8.991815706465342, 0.5156084030180363, 3.703827688886369},
		},
	}

	for _, c := range cases {
		actual := c.color.Lab()

		assert.InDelta(t, c.expected.L, actual.L, 0.005)
		assert.InDelta(t, c.expected.A, actual.A, 0.005)
		assert.InDelta(t, c.expected.B, actual.B, 0.005)
	}
}
