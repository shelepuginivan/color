package color_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestLab(t *testing.T) {
	assert.Implements(t, (*color.Color)(nil), color.Lab{})
}

func TestLab_String(t *testing.T) {
	for range 1000 {
		var (
			l = rand.Float64() * 100
			a = rand.Float64() * 100
			b = rand.Float64() * 100
		)

		var (
			expected = fmt.Sprintf("lab(%.4f, %.4f, %.4f)", l, a, b)
			actual   = color.NewLab(l, a, b).String()
		)

		assert.Equal(t, expected, actual)
	}
}

func TestLab_Lch(t *testing.T) {
	cases := []struct {
		color    *color.Lab
		expected *color.Lch
	}{
		{
			&color.Lab{8.991815706465342, 0.5156084030180363, 3.703827688886369},
			&color.Lch{8.991815706465342, 3.7396951758251333, 82},
		},
	}

	for _, c := range cases {
		actual := c.color.Lch()

		assert.InDelta(t, c.expected.L, actual.L, 0.001)
		assert.InDelta(t, c.expected.C, actual.C, 0.001)
		assert.Equal(t, c.expected.H, actual.H)
	}
}
