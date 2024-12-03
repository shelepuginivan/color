package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestLab(t *testing.T) {
	assert.Implements(t, (*interface {
		CMYK() *color.CMYK
		Hex() string
		HSL() *color.HSL
		HSV() *color.HSV
		RGB() *color.RGB
		XYZ() *color.XYZ

		String() string
	})(nil), color.Lab{})
}
