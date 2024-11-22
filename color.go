package color

import (
	"fmt"
	"strconv"
)

// Color represents a color.
type Color struct {
	R uint8 // Red.
	G uint8 // Green.
	B uint8 // Blue.
}

// New returns a new instance of [Color].
func New(r, g, b uint8) *Color {
	return &Color{r, g, b}
}

// NewFromHex returns a new instance of [Color] by parsing hexadecimal color
// string. The string may start with hash character (`#`) and may be either
// short or long hexadecimal color. E.g. `fff`, `#fff`, `ffffff`, `#ffffff` are
// all valid arguments.
func NewFromHex(hex string) (*Color, error) {
	i := 0
	if hex[0] == '#' {
		i = 1
	}

	var hex_r, hex_g, hex_b string

	switch len(hex) - i {
	case 3: // Short hexadecimal notation, e.g. `#abc`.
		hex_r = string([]byte{hex[i], hex[i]})
		hex_g = string([]byte{hex[i+1], hex[i+1]})
		hex_b = string([]byte{hex[i+2], hex[i+2]})
	case 6: // Long hexadecimal notation, e.g. `#aabbcc`
		hex_r = hex[i : i+2]
		hex_g = hex[i+2 : i+4]
		hex_b = hex[i+4 : i+6]
	default:
		return nil, fmt.Errorf("invalid hexadecimal string")
	}

	r, err := strconv.ParseUint(hex_r, 16, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid value of red channel: %v", err)
	}

	g, err := strconv.ParseUint(hex_g, 16, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid value of green channel: %v", err)
	}

	b, err := strconv.ParseUint(hex_b, 16, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid value of blue channel: %v", err)
	}

	return &Color{uint8(r), uint8(g), uint8(b)}, nil
}

func (c Color) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}
