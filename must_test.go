package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	assert.NotPanics(t, func() {
		color.Must(color.NewFromHex("#ffffff"))
	})

	assert.Panics(t, func() {
		color.Must(color.NewFromHex(":("))
	})
}
