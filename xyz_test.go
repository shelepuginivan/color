package color_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestXYZ(t *testing.T) {
	assert.Implements(t, (*color.Color)(nil), color.XYZ{})
}

func TestXYZ_Lab(t *testing.T) {
	cases := []struct {
		color    *color.XYZ
		expected *color.Lab
	}{
		{&color.XYZ{0, 0, 0}, &color.Lab{0, 0, 0}},
		{&color.XYZ{0.41246, 0.21267, 0.01933}, &color.Lab{53.27, 80.109, 67.220}},
		{&color.XYZ{0.35758, 0.71515, 0.11919}, &color.Lab{87.73, -86.1846, 83.181}},
		{&color.XYZ{0.18044, 0.07217, 0.95030}, &color.Lab{32.30, 79.19667, -107.86368}},
		{&color.XYZ{0.77003, 0.92783, 0.13853}, &color.Lab{97.13, -21.5559, 94.4845}},
		{&color.XYZ{0.59289, 0.28485, 0.96964}, &color.Lab{60.3199, 98.2542, -60.84298}},
		{&color.XYZ{0.53801, 0.78733, 1.06950}, &color.Lab{91.1165, -48.0796, -14.1381}},
		{&color.XYZ{0.95047, 1.00000, 1.08883}, &color.Lab{100.0, 0.00526, -0.0104}},
	}

	for _, c := range cases {
		actual := c.color.Lab()

		assert.InDelta(t, c.expected.L, actual.L, 0.05)
		assert.InDelta(t, c.expected.A, actual.A, 0.05)
		assert.InDelta(t, c.expected.B, actual.B, 0.05)
	}
}

func TestXYZ_LabWithReferenceWhite(t *testing.T) {
	cases := []struct {
		color    *color.XYZ
		white    *color.XYZ
		expected *color.Lab
	}{
		{&color.XYZ{0, 0, 0}, color.D65, &color.Lab{0, 0, 0}},
		{&color.XYZ{0.41246, 0.21267, 0.01933}, color.D65, &color.Lab{53.27, 80.109, 67.220}},
		{&color.XYZ{0.35758, 0.71515, 0.11919}, color.D65, &color.Lab{87.73, -86.1846, 83.181}},
		{&color.XYZ{0.18044, 0.07217, 0.95030}, color.D65, &color.Lab{32.30, 79.19667, -107.86368}},
		{&color.XYZ{0.77003, 0.92783, 0.13853}, color.D65, &color.Lab{97.13, -21.5559, 94.4845}},
		{&color.XYZ{0.59289, 0.28485, 0.96964}, color.D65, &color.Lab{60.3199, 98.2542, -60.84298}},
		{&color.XYZ{0.53801, 0.78733, 1.06950}, color.D65, &color.Lab{91.1165, -48.0796, -14.1381}},
		{&color.XYZ{0.95047, 1.000, 1.08883}, color.D65, &color.Lab{100.0, 0.00526, -0.0104}},
	}

	for _, c := range cases {
		actual := c.color.LabWithWhitepoint(c.white)

		assert.InDelta(t, c.expected.L, actual.L, 0.05)
		assert.InDelta(t, c.expected.A, actual.A, 0.05)
		assert.InDelta(t, c.expected.B, actual.B, 0.05)
	}
}

func TestXYZ_Oklab(t *testing.T) {
	cases := []struct {
		color    *color.XYZ
		expected *color.Oklab
	}{
		{&color.XYZ{0.20654008, 0.12197225, 0.05136952}, &color.Oklab{0.5163401, 0.154695, 0.0628957}},
	}

	for _, c := range cases {
		actual := c.color.Oklab()

		assert.InDelta(t, c.expected.L, actual.L, 0.05)
		assert.InDelta(t, c.expected.A, actual.A, 0.05)
		assert.InDelta(t, c.expected.B, actual.B, 0.05)
	}
}

func TestXYZ_RGB(t *testing.T) {
	cases := []struct {
		color    *color.XYZ
		expected *color.RGB
	}{
		{color.D65, &color.RGB{255, 255, 255}},
		{&color.XYZ{0.41246, 0.21267, 0.01933}, &color.RGB{255, 0, 0}},
		{&color.XYZ{0.35758, 0.71515, 0.11919}, &color.RGB{0, 255, 0}},
		{&color.XYZ{0.18044, 0.07217, 0.95030}, &color.RGB{0, 0, 255}},
		{&color.XYZ{0.77003, 0.92783, 0.13853}, &color.RGB{255, 255, 0}},
		{&color.XYZ{0.59289, 0.28485, 0.96964}, &color.RGB{255, 0, 255}},
		{&color.XYZ{0.53801, 0.78733, 1.06950}, &color.RGB{0, 255, 255}},
		{&color.XYZ{0.95047, 1.00000, 1.08883}, &color.RGB{255, 255, 255}},
		{&color.XYZ{0.21461, 0.30457, 0.18932}, &color.RGB{102, 164, 108}},
		{&color.XYZ{0.26232, 0.19865, 0.34794}, &color.RGB{164, 102, 158}},
		{&color.XYZ{0.24425, 0.22286, 0.15583}, &color.RGB{164, 120, 102}},
		{&color.XYZ{0.84203, 0.88676, 0.97457}, &color.RGB{241, 242, 243}},
		{&color.XYZ{0.46053, 0.38455, 0.09271}, &color.RGB{238, 144, 60}},
	}

	for _, c := range cases {
		actual := c.color.RGB()
		assert.Equal(t, c.expected, actual)
	}
}

func TestXYZ_String(t *testing.T) {
	for range 1000 {
		var (
			x = rand.Float64() * 100
			y = rand.Float64() * 100
			z = rand.Float64() * 100
		)

		var (
			expected = fmt.Sprintf("xyz(%.4f, %.4f, %.4f)", x, y, z)
			actual   = color.NewXYZ(x, y, z).String()
		)

		assert.Equal(t, expected, actual)
	}
}
