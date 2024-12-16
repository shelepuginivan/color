package color_test

import (
	"fmt"
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	assert.NotPanics(t, func() {
		color.Must(color.ParseHex("#ffffff"))
	})

	assert.NotPanics(t, func() {
		color.Must(color.RGB{}, nil)
	})

	assert.Panics(t, func() {
		color.Must(color.ParseHex(":("))
	})

	assert.Panics(t, func() {
		color.Must(color.RGB{}, fmt.Errorf("should panic"))
	})

	assert.Equal(t, color.HSL{120, 50, 50}, color.Must(color.HSL{120, 50, 50}, nil))
}
