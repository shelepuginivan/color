package normalize_test

import (
	"testing"

	"github.com/shelepuginivan/color/internal/normalize"
	"github.com/stretchr/testify/assert"
)

func TestPercents(t *testing.T) {
	cases := []struct {
		value    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{-1, 0},
		{50, 50},
		{101, 100},
		{1000000000, 100},
	}

	for _, c := range cases {
		actual := normalize.Percents(c.value)
		assert.Equal(t, c.expected, actual)
	}
}

func TestPercentsFloat(t *testing.T) {
	cases := []struct {
		value    int
		expected float64
	}{
		{-1, 0},
		{0, 0},
		{29, 0.29},
		{21, 0.21},
		{3, 0.03},
		{95, 0.95},
		{100, 1},
		{101, 1},
		{1000000, 1},
	}

	for _, c := range cases {
		actual := normalize.PercentsFloat(c.value)
		assert.Equal(t, c.expected, actual)
	}
}

func TestFloatPercents(t *testing.T) {
	cases := []struct {
		value    float64
		expected int
	}{
		{-1, 0},
		{0, 0},
		{0.01, 1},
		{0.98, 98},
		{0.63, 63},
		{0.2, 20},
		{1, 100},
		{2, 100},
		{10000, 100},
	}

	for _, c := range cases {
		actual := normalize.FloatPercents(c.value)
		assert.Equal(t, c.expected, actual)
	}
}

func TestDegrees(t *testing.T) {
	cases := []struct {
		value    int
		expected int
	}{
		{0, 0},
		{10, 10},
		{200, 200},
		{360, 0},
		{720, 0},
		{500, 140},
		{-120, 240},
		{-1, 359},
	}

	for _, c := range cases {
		actual := normalize.Degrees(c.value)
		assert.Equal(t, c.expected, actual)
	}
}

func TestDegreesFloat(t *testing.T) {
	cases := []struct {
		value    int
		expected float64
	}{
		{0, 0},
		{10, float64(10) / 360},
		{200, float64(200) / 360},
		{360, float64(0) / 360},
		{720, float64(0) / 360},
		{500, float64(140) / 360},
		{-120, float64(240) / 360},
		{-1, float64(359) / 360},
	}

	for _, c := range cases {
		actual := normalize.DegreesFloat(c.value)
		assert.Equal(t, c.expected, actual)
	}
}
