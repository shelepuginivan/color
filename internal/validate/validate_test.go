package validate_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/shelepuginivan/color/internal/validate"
	"github.com/stretchr/testify/assert"
)

func TestIsPercent(t *testing.T) {
	cases := []struct {
		v        int
		a        string
		expected error
	}{
		{0, "arg", nil},
		{100, "arg", nil},
		{50, "arg", nil},
		{69, "arg", nil},
		{-1, "arg", fmt.Errorf("arg must be a valid value in percents (integer in range [0, 100]), got -1")},
		{101, "b", fmt.Errorf("b must be a valid value in percents (integer in range [0, 100]), got 101")},
	}

	for _, c := range cases {
		actual := validate.IsPercent(c.v, c.a)
		assert.Equal(t, c.expected, actual)
	}

	for range 1000 {
		p := rand.IntN(1000) - 500
		err := validate.IsPercent(p, "")

		if 0 <= p && p <= 100 {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestIsDegree(t *testing.T) {
	cases := []struct {
		v        int
		a        string
		expected error
	}{
		{0, "arg", nil},
		{360, "arg", nil},
		{180, "arg", nil},
		{277, "arg", nil},
		{-1, "arg", fmt.Errorf("arg must be a valid value in degrees (integer in range [0, 360]), got -1")},
		{361, "deg", fmt.Errorf("deg must be a valid value in degrees (integer in range [0, 360]), got 361")},
	}

	for _, c := range cases {
		actual := validate.IsDegree(c.v, c.a)
		assert.Equal(t, c.expected, actual)
	}

	for range 1000 {
		d := rand.IntN(1000) - 500
		err := validate.IsDegree(d, "")

		if 0 <= d && d <= 360 {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestAll(t *testing.T) {
	cases := []struct {
		errs     []error
		expected error
	}{
		{
			errs:     []error{nil, nil, nil, nil, nil},
			expected: nil,
		},
		{
			errs:     []error{},
			expected: nil,
		},
		{
			errs:     []error{nil, fmt.Errorf("one"), nil, fmt.Errorf("two"), fmt.Errorf("three")},
			expected: fmt.Errorf("one"),
		},
	}

	for _, c := range cases {
		actual := validate.All(c.errs...)
		assert.Equal(t, c.expected, actual)
	}
}
