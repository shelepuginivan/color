package color

import (
	"fmt"
	"math"
	"strconv"
)

// Percent convertation constant.
const percent = 100

// Color represents a color.
type Color struct {
	R uint8 // Red.
	G uint8 // Green.
	B uint8 // Blue.
}

// HSL representation of color.
type HSL struct {
	Hue        int // Hue (degrees).
	Saturation int // Saturation (percent).
	Lightness  int // Lightness (percent).
}

// CMYK representation of color.
type CMYK struct {
	C int // Cyan (percent).
	M int // Magenta (percent).
	Y int // Yellow (percent).
	K int // Black key (percent).
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

func NewFromCMYK(cyan, magenta, yellow, key int) (*Color, error) {
	validity := [...]error{
		isPercent(cyan, "cyan"),
		isPercent(magenta, "magenta"),
		isPercent(yellow, "yellow"),
		isPercent(key, "key"),
	}

	for _, err := range validity {
		if err != nil {
			return nil, err
		}
	}

	c := float64(cyan) / 100
	m := float64(magenta) / 100
	y := float64(yellow) / 100
	k := float64(key) / 100

	r := 255 * (1 - c) * (1 - k)
	g := 255 * (1 - m) * (1 - k)
	b := 255 * (1 - y) * (1 - k)

	return &Color{
		R: uint8(math.Round(r)),
		G: uint8(math.Round(g)),
		B: uint8(math.Round(b)),
	}, nil
}

// NewFromHSL returns a new instance of [Color] by converting HSL to RGB.
func NewFromHSL(hue, saturation, lightness int) (*Color, error) {
	validity := [...]error{
		isDegree(hue, "hue"),
		isPercent(saturation, "saturation"),
		isPercent(lightness, "lightness"),
	}

	for _, err := range validity {
		if err != nil {
			return nil, err
		}
	}

	h := float64(hue) / 360
	s := float64(saturation) / 100
	l := float64(lightness) / 100

	// The default case is when the color is achromatic.
	var (
		r = l
		g = l
		b = l
	)

	if s != 0 {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		r = hueToRGB(p, q, h+1.0/3)
		g = hueToRGB(p, q, h)
		b = hueToRGB(p, q, h-1.0/3)
	}

	return &Color{
		R: uint8(math.Round(r * 255)),
		G: uint8(math.Round(g * 255)),
		B: uint8(math.Round(b * 255)),
	}, nil
}

// Hex returns hexadecimal representation of color.
func (c Color) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// CMYK returns [CMYK] representation of color (cyan, magenta, yellow, key).
func (c Color) CMYK() *CMYK {
	r := float64(c.R) / 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255

	key := 1 - max(r, g, b)
	d := 1 - key

	// The default case is when key equals 1, i.e. the color is black.
	var (
		cyan    = 0.0
		magenta = 0.0
		yellow  = 0.0
	)

	if d != 0 {
		cyan = (d - r) / d
		magenta = (d - g) / d
		yellow = (d - b) / d
	}

	return &CMYK{
		C: int(math.Round(cyan * 100)),
		M: int(math.Round(magenta * 100)),
		Y: int(math.Round(yellow * 100)),
		K: int(math.Round(key * 100)),
	}
}

// HSL returns [HSL] representation of color (hue, saturation, lightness).
func (c Color) HSL() *HSL {
	r := float64(c.R) / 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255

	mx := max(r, g, b)
	mn := min(r, g, b)
	h, s, l := 0.0, 0.0, (mx+mn)/2

	d := mx - mn
	if d != 0 {
		switch mx {
		case r:
			h = (g - b) / d
		case g:
			h = (b-r)/d + 2
		default:
			h = (r-g)/d + 4
		}

		h *= 60
		if h < 0 {
			h += 360
		}
		s = d / (1 - max(2*l-1, 1-2*l))
	}

	return &HSL{
		Hue:        int(h),
		Saturation: int(s * 100),
		Lightness:  int(l * 100),
	}
}

// hueToRGB calculates the RGB value for a given hue component. It takes three
// parameters: p and q are the intermediate values calculated from the HSL
// representation, and t represents the normalized hue value (ranging from 0 to
// 1). The function adjusts the value of t to ensure it falls within the range
// [0, 1] by wrapping around if necessary. It then computes the RGB value based
// on the hue's position in the color wheel, returning the corresponding
// float64 value for the red, green, or blue component. This function is used
// internally in the HSL to RGB conversion process.
func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6 {
		return p + (q-p)*6*t
	}
	if t < 0.5 {
		return q
	}
	if t < 2.0/3 {
		return p + (q-p)*(2.0/3-t)*6
	}
	return p
}

// isPercent validates value in percents.
func isPercent(p int, argname string) error {
	if p < 0 || p > 100 {
		return fmt.Errorf("%s must be a valid value in percents (integer in range [0, 100]), got %d", argname, p)
	}
	return nil
}

// isDegree validates value in degrees.
func isDegree(d int, argname string) error {
	if d < 0 || d > 360 {
		return fmt.Errorf("%s must be a valid value in degress (integer in range [0, 360]), got %d", argname, d)
	}
	return nil
}
