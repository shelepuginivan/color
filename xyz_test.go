package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestXYZ(t *testing.T) {
	assert.Implements(t, (*interface {
		CMYK() *color.CMYK
		Lab() *color.Lab
		RGB() *color.RGB
		Hex() string
		HSL() *color.HSL
		HSV() *color.HSV

		String() string
	})(nil), color.XYZ{})
}

func TestXYZ_Lab(t *testing.T) {
	defaultWhite := &color.XYZ{
		X: color.ReferenceWhiteX,
		Y: color.ReferenceWhiteY,
		Z: color.ReferenceWhiteZ,
	}

	cases := []struct {
		color    *color.XYZ
		expected *color.Lab
	}{
		{&color.XYZ{0, 0, 0}, &color.Lab{0, 0, 0, defaultWhite}},
		{&color.XYZ{41.246, 21.267, 1.933}, &color.Lab{53.27, 80.109, 67.220, defaultWhite}},
		{&color.XYZ{35.758, 71.515, 11.919}, &color.Lab{87.73, -86.1846, 83.181, defaultWhite}},
		{&color.XYZ{18.044, 7.217, 95.030}, &color.Lab{32.30, 79.19667, -107.86368, defaultWhite}},
		{&color.XYZ{77.003, 92.783, 13.853}, &color.Lab{97.13, -21.5559, 94.4845, defaultWhite}},
		{&color.XYZ{59.289, 28.485, 96.964}, &color.Lab{60.3199, 98.2542, -60.84298, defaultWhite}},
		{&color.XYZ{53.801, 78.733, 106.950}, &color.Lab{91.1165, -48.0796, -14.1381, defaultWhite}},
		{&color.XYZ{95.047, 100.000, 108.883}, &color.Lab{100.0, 0.00526, -0.0104, defaultWhite}},
	}

	for _, c := range cases {
		actual := c.color.Lab()

		assert.InDelta(t, c.expected.L, actual.L, 0.05)
		assert.InDelta(t, c.expected.A, actual.A, 0.05)
		assert.InDelta(t, c.expected.B, actual.B, 0.05)
	}
}

func TestXYZ_LabWithReferenceWhite(t *testing.T) {
	defaultWhite := &color.XYZ{
		X: color.ReferenceWhiteX,
		Y: color.ReferenceWhiteY,
		Z: color.ReferenceWhiteZ,
	}

	cases := []struct {
		color    *color.XYZ
		white    *color.XYZ
		expected *color.Lab
	}{
		{&color.XYZ{0, 0, 0}, defaultWhite, &color.Lab{0, 0, 0, defaultWhite}},
		{&color.XYZ{41.246, 21.267, 1.933}, defaultWhite, &color.Lab{53.27, 80.109, 67.220, defaultWhite}},
		{&color.XYZ{35.758, 71.515, 11.919}, defaultWhite, &color.Lab{87.73, -86.1846, 83.181, defaultWhite}},
		{&color.XYZ{18.044, 7.217, 95.030}, defaultWhite, &color.Lab{32.30, 79.19667, -107.86368, defaultWhite}},
		{&color.XYZ{77.003, 92.783, 13.853}, defaultWhite, &color.Lab{97.13, -21.5559, 94.4845, defaultWhite}},
		{&color.XYZ{59.289, 28.485, 96.964}, defaultWhite, &color.Lab{60.3199, 98.2542, -60.84298, defaultWhite}},
		{&color.XYZ{53.801, 78.733, 106.950}, defaultWhite, &color.Lab{91.1165, -48.0796, -14.1381, defaultWhite}},
		{&color.XYZ{95.047, 100.000, 108.883}, defaultWhite, &color.Lab{100.0, 0.00526, -0.0104, defaultWhite}},
	}

	for _, c := range cases {
		actual := c.color.LabWithReferenceWhite(c.white)

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
		{color.DefaultReferenceWhite, &color.RGB{255, 255, 255}},
		{&color.XYZ{41.246, 21.267, 1.933}, &color.RGB{255, 0, 0}},
		{&color.XYZ{35.758, 71.515, 11.919}, &color.RGB{0, 255, 0}},
		{&color.XYZ{18.044, 7.217, 95.030}, &color.RGB{0, 0, 255}},
		{&color.XYZ{77.003, 92.783, 13.853}, &color.RGB{255, 255, 0}},
		{&color.XYZ{59.289, 28.485, 96.964}, &color.RGB{255, 0, 255}},
		{&color.XYZ{53.801, 78.733, 106.950}, &color.RGB{0, 255, 255}},
		{&color.XYZ{95.047, 100.000, 108.883}, &color.RGB{255, 255, 255}},
		{&color.XYZ{21.461, 30.457, 18.932}, &color.RGB{102, 164, 108}},
		{&color.XYZ{26.232, 19.865, 34.794}, &color.RGB{164, 102, 158}},
		{&color.XYZ{24.425, 22.286, 15.583}, &color.RGB{164, 120, 102}},
		{&color.XYZ{84.203, 88.676, 97.457}, &color.RGB{241, 242, 243}},
		{&color.XYZ{46.053, 38.455, 9.271}, &color.RGB{238, 144, 60}},
	}

	for _, c := range cases {
		actual := c.color.RGB()
		assert.Equal(t, c.expected, actual)
	}
}
