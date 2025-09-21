package color_test

import (
	"testing"

	"github.com/shelepuginivan/color"
	"github.com/stretchr/testify/assert"
)

func TestOklab(t *testing.T) {
	assert.Implements(t, (*color.Color)(nil), color.Oklab{})
}
