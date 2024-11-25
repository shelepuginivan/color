package color_test

import (
	"fmt"
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

func ExampleMust() {
	yellow := color.Must(color.NewFromCMYK(0, 0, 100, 0))

	fmt.Println(yellow.Hex()) // Output: #ffff00
}
