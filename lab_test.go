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
