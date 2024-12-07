package degrees_test

import (
	"math"
	"testing"

	"github.com/shelepuginivan/color/internal/degrees"
	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	cases := []struct {
		v        int
		expected int
	}{
		{0, 0},
		{360, 0},
		{720, 0},
		{-360, 0},
		{-1, 359},
		{370, 10},
		{450, 90},
		{-450, 270},
	}

	for _, c := range cases {
		actual := degrees.Normalize(c.v)
		assert.Equal(t, c.expected, actual)
	}
}

func TestToFloat(t *testing.T) {
	cases := []struct {
		v        int
		expected float64
	}{
		{0, 0.0},
		{360, 0.0},
		{180, 0.5},
		{90, 0.25},
		{-90, 0.75},
		{450, 0.25},
	}

	for _, c := range cases {
		actual := degrees.ToFloat(c.v)
		assert.Equal(t, c.expected, actual)
	}
}

func TestFromFloat(t *testing.T) {
	cases := []struct {
		v        float64
		expected int
	}{
		{0.0, 0},
		{0.5, 180},
		{0.25, 90},
		{0.75, 270},
		{1.0, 0},
		{0.1, 36},
	}

	for _, c := range cases {
		actual := degrees.FromFloat(c.v)
		assert.Equal(t, c.expected, actual)
	}
}

func TestToRadians(t *testing.T) {
	cases := []struct {
		input    int
		expected float64
	}{
		{0, 0.0},
		{180, math.Pi},
		{90, math.Pi / 2},
		{360, 0.0},
		{270, 3 * math.Pi / 2},
		{-90, 3 * math.Pi / 2}, // +360
	}

	for _, c := range cases {
		actual := degrees.ToRadians(c.input)
		assert.Equal(t, c.expected, actual)
	}
}

func TestFromRadians(t *testing.T) {
	cases := []struct {
		input    float64
		expected int
	}{
		{0.0, 0},
		{math.Pi, 180},
		{math.Pi / 2, 90},
		{-math.Pi / 2, 270},
	}

	for _, c := range cases {
		actual := degrees.FromRadians(c.input)
		assert.Equal(t, c.expected, actual)
	}
}
