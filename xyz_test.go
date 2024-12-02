package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestXYZ_Lab(t *testing.T) {
	cases := []struct {
		color    *color.XYZ
		expected *color.Lab
	}{
		{&color.XYZ{0, 0, 0}, &color.Lab{0, 0, 0}},
		{&color.XYZ{41.246, 21.267, 1.933}, &color.Lab{53.27, 80.109, 67.220}},
		{&color.XYZ{35.758, 71.515, 11.919}, &color.Lab{87.73, -86.1846, 83.181}},
		{&color.XYZ{18.044, 7.217, 95.030}, &color.Lab{32.30, 79.19667, -107.86368}},
		{&color.XYZ{77.003, 92.783, 13.853}, &color.Lab{97.13, -21.5559, 94.4845}},
		{&color.XYZ{59.289, 28.485, 96.964}, &color.Lab{60.3199, 98.2542, -60.84298}},
		{&color.XYZ{53.801, 78.733, 106.950}, &color.Lab{91.1165, -48.0796, -14.1381}},
		{&color.XYZ{95.047, 100.000, 108.883}, &color.Lab{100.0, 0.00526, -0.0104}},
	}

	for _, c := range cases {
		actual := c.color.Lab()

		assert.InDelta(t, c.expected.L, actual.L, 0.05)
		assert.InDelta(t, c.expected.A, actual.A, 0.05)
		assert.InDelta(t, c.expected.B, actual.B, 0.05)
	}
}
