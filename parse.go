package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shelepuginivan/color/internal/degrees"
	"github.com/shelepuginivan/color/internal/percents"
)

// Parse is a generic color parsing function.
//
// The following order is used for parsing:
//  1. [CSS named colors] using [color.ParseNamed]
//  2. Hexadecimal notation using [color.ParseHex]
//  3. Color functions using [color.ParseFunc]
//
// See [color.ParseNamed], [color.ParseHex], and [color.ParseFunc] for more
// information about color notations supported by Parse.
//
// [CSS named colors]: https://developer.mozilla.org/en-US/docs/Web/CSS/named-color
func Parse(s string) (Color, error) {
	c, err := ParseNamed(s)
	if err == nil {
		return c, nil
	}

	c, err = ParseHex(s)
	if err == nil {
		return c, nil
	}

	c, err = ParseFunc(s)
	if err == nil {
		return c, nil
	}

	return nil, fmt.Errorf("failed to parse input string")
}

// namedColorMap maps [CSS named colors] to [RGB]. See [Reference].
//
// [CSS named colors]: https://developer.mozilla.org/en-US/docs/Web/CSS/named-color
// [Reference]: https://www.rapidtables.com/web/color/RGB_Color.html
var namedColorMap = map[string]RGB{
	"maroon":               {128, 0, 0},     // #800000
	"darkred":              {139, 0, 0},     // #8b0000
	"brown":                {165, 42, 42},   // #a52a2a
	"firebrick":            {178, 34, 34},   // #b22222
	"crimson":              {220, 20, 60},   // #dc143c
	"red":                  {255, 0, 0},     // #ff0000
	"tomato":               {255, 99, 71},   // #ff6347
	"coral":                {255, 127, 80},  // #ff7f50
	"indianred":            {205, 92, 92},   // #cd5c5c
	"lightcoral":           {240, 128, 128}, // #f08080
	"darksalmon":           {233, 150, 122}, // #e9967a
	"salmon":               {250, 128, 114}, // #fa8072
	"lightsalmon":          {255, 160, 122}, // #ffa07a
	"orangered":            {255, 69, 0},    // #ff4500
	"darkorange":           {255, 140, 0},   // #ff8c00
	"orange":               {255, 165, 0},   // #ffa500
	"gold":                 {255, 215, 0},   // #ffd700
	"darkgoldenrod":        {184, 134, 11},  // #b8860b
	"goldenrod":            {218, 165, 32},  // #daa520
	"palegoldenrod":        {238, 232, 170}, // #eee8aa
	"darkkhaki":            {189, 183, 107}, // #bdb76b
	"khaki":                {240, 230, 140}, // #f0e68c
	"olive":                {128, 128, 0},   // #808000
	"yellow":               {255, 255, 0},   // #ffff00
	"yellowgreen":          {154, 205, 50},  // #9acd32
	"darkolivegreen":       {85, 107, 47},   // #556b2f
	"olivedrab":            {107, 142, 35},  // #6b8e23
	"lawngreen":            {124, 252, 0},   // #7cfc00
	"chartreuse":           {127, 255, 0},   // #7fff00
	"greenyellow":          {173, 255, 47},  // #adff2f
	"darkgreen":            {0, 100, 0},     // #006400
	"green":                {0, 128, 0},     // #008000
	"forestgreen":          {34, 139, 34},   // #228b22
	"lime":                 {0, 255, 0},     // #00ff00
	"limegreen":            {50, 205, 50},   // #32cd32
	"lightgreen":           {144, 238, 144}, // #90ee90
	"palegreen":            {152, 251, 152}, // #98fb98
	"darkseagreen":         {143, 188, 143}, // #8fbc8f
	"mediumspringgreen":    {0, 250, 154},   // #00fa9a
	"springgreen":          {0, 255, 127},   // #00ff7f
	"seagreen":             {46, 139, 87},   // #2e8b57
	"mediumaquamarine":     {102, 205, 170}, // #66cdaa
	"mediumseagreen":       {60, 179, 113},  // #3cb371
	"lightseagreen":        {32, 178, 170},  // #20b2aa
	"darkslategray":        {47, 79, 79},    // #2f4f4f
	"teal":                 {0, 128, 128},   // #008080
	"darkcyan":             {0, 139, 139},   // #008b8b
	"aqua":                 {0, 255, 255},   // #00ffff
	"cyan":                 {0, 255, 255},   // #00ffff
	"lightcyan":            {224, 255, 255}, // #e0ffff
	"darkturquoise":        {0, 206, 209},   // #00ced1
	"turquoise":            {64, 224, 208},  // #40e0d0
	"mediumturquoise":      {72, 209, 204},  // #48d1cc
	"paleturquoise":        {175, 238, 238}, // #afeeee
	"aquamarine":           {127, 255, 212}, // #7fffd4
	"powderblue":           {176, 224, 230}, // #b0e0e6
	"cadetblue":            {95, 158, 160},  // #5f9ea0
	"steelblue":            {70, 130, 180},  // #4682b4
	"cornflowerblue":       {100, 149, 237}, // #6495ed
	"deepskyblue":          {0, 191, 255},   // #00bfff
	"dodgerblue":           {30, 144, 255},  // #1e90ff
	"lightblue":            {173, 216, 230}, // #add8e6
	"skyblue":              {135, 206, 235}, // #87ceeb
	"lightskyblue":         {135, 206, 250}, // #87cefa
	"midnightblue":         {25, 25, 112},   // #191970
	"navy":                 {0, 0, 128},     // #000080
	"darkblue":             {0, 0, 139},     // #00008b
	"mediumblue":           {0, 0, 205},     // #0000cd
	"blue":                 {0, 0, 255},     // #0000ff
	"royalblue":            {65, 105, 225},  // #4169e1
	"blueviolet":           {138, 43, 226},  // #8a2be2
	"indigo":               {75, 0, 130},    // #4b0082
	"darkslateblue":        {72, 61, 139},   // #483d8b
	"slateblue":            {106, 90, 205},  // #6a5acd
	"mediumslateblue":      {123, 104, 238}, // #7b68ee
	"mediumpurple":         {147, 112, 219}, // #9370db
	"darkmagenta":          {139, 0, 139},   // #8b008b
	"darkviolet":           {148, 0, 211},   // #9400d3
	"darkorchid":           {153, 50, 204},  // #9932cc
	"mediumorchid":         {186, 85, 211},  // #ba55d3
	"purple":               {128, 0, 128},   // #800080
	"thistle":              {216, 191, 216}, // #d8bfd8
	"plum":                 {221, 160, 221}, // #dda0dd
	"violet":               {238, 130, 238}, // #ee82ee
	"magenta":              {255, 0, 255},   // #ff00ff
	"fuchsia":              {255, 0, 255},   // #ff00ff
	"orchid":               {218, 112, 214}, // #da70d6
	"mediumvioletred":      {199, 21, 133},  // #c71585
	"palevioletred":        {219, 112, 147}, // #db7093
	"deeppink":             {255, 20, 147},  // #ff1493
	"hotpink":              {255, 105, 180}, // #ff69b4
	"lightpink":            {255, 182, 193}, // #ffb6c1
	"pink":                 {255, 192, 203}, // #ffc0cb
	"antiquewhite":         {250, 235, 215}, // #faebd7
	"beige":                {245, 245, 220}, // #f5f5dc
	"bisque":               {255, 228, 196}, // #ffe4c4
	"blanchedalmond":       {255, 235, 205}, // #ffebcd
	"wheat":                {245, 222, 179}, // #f5deb3
	"cornsilk":             {255, 248, 220}, // #fff8dc
	"lemonchiffon":         {255, 250, 205}, // #fffacd
	"lightgoldenrodyellow": {250, 250, 210}, // #fafad2
	"lightyellow":          {255, 255, 224}, // #ffffe0
	"saddlebrown":          {139, 69, 19},   // #8b4513
	"sienna":               {160, 82, 45},   // #a0522d
	"chocolate":            {210, 105, 30},  // #d2691e
	"peru":                 {205, 133, 63},  // #cd853f
	"sandybrown":           {244, 164, 96},  // #f4a460
	"burlywood":            {222, 184, 135}, // #deb887
	"tan":                  {210, 180, 140}, // #d2b48c
	"rosybrown":            {188, 143, 143}, // #bc8f8f
	"moccasin":             {255, 228, 181}, // #ffe4b5
	"navajowhite":          {255, 222, 173}, // #ffdead
	"peachpuff":            {255, 218, 185}, // #ffdab9
	"mistyrose":            {255, 228, 225}, // #ffe4e1
	"lavenderblush":        {255, 240, 245}, // #fff0f5
	"linen":                {250, 240, 230}, // #faf0e6
	"oldlace":              {253, 245, 230}, // #fdf5e6
	"papayawhip":           {255, 239, 213}, // #ffefd5
	"seashell":             {255, 245, 238}, // #fff5ee
	"mintcream":            {245, 255, 250}, // #f5fffa
	"slategray":            {112, 128, 144}, // #708090
	"lightslategray":       {119, 136, 153}, // #778899
	"lightsteelblue":       {176, 196, 222}, // #b0c4de
	"lavender":             {230, 230, 250}, // #e6e6fa
	"floralwhite":          {255, 250, 240}, // #fffaf0
	"aliceblue":            {240, 248, 255}, // #f0f8ff
	"ghostwhite":           {248, 248, 255}, // #f8f8ff
	"honeydew":             {240, 255, 240}, // #f0fff0
	"ivory":                {255, 255, 240}, // #fffff0
	"azure":                {240, 255, 255}, // #f0ffff
	"snow":                 {255, 250, 250}, // #fffafa
	"black":                {0, 0, 0},       // #000000
	"dimgray":              {105, 105, 105}, // #696969
	"dimgrey":              {105, 105, 105}, // #696969
	"gray":                 {128, 128, 128}, // #808080
	"grey":                 {128, 128, 128}, // #808080
	"darkgray":             {169, 169, 169}, // #a9a9a9
	"darkgrey":             {169, 169, 169}, // #a9a9a9
	"silver":               {192, 192, 192}, // #c0c0c0
	"lightgray":            {211, 211, 211}, // #d3d3d3
	"lightgrey":            {211, 211, 211}, // #d3d3d3
	"gainsboro":            {220, 220, 220}, // #dcdcdc
	"whitesmoke":           {245, 245, 245}, // #f5f5f5
	"white":                {255, 255, 255}, // #ffffff
}

// ParseNamed returns a [CSS named color] by name.
//
// [CSS named color]: https://developer.mozilla.org/en-US/docs/Web/CSS/named-color
func ParseNamed(name string) (Color, error) {
	c, ok := namedColorMap[strings.ToLower(name)]
	if !ok {
		return nil, fmt.Errorf("unknown named color: %s", name)
	}
	return &c, nil
}

// ParseHex returns a color by parsing hexadecimal color string. The string may
// start with hash character (`#`) and may be either short or long hexadecimal
// color. Hence, all of the following strings are valid arguments:
//   - fff
//   - #fff
//   - ffffff
//   - #ffffff
func ParseHex(hex string) (Color, error) {
	i := 0
	if strings.HasPrefix(hex, "#") {
		i++
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

	return &RGB{uint8(r), uint8(g), uint8(b)}, nil
}

// ParseFunc parses the color function and returns a color.
//
// A color function is a string starting with the function name (e.g. rgb),
// containing its parameters in parentheses. Some examples are:
//   - rgb(255, 255, 255)
//   - hsl(0.8turn 80% 30%)
//   - xyz(0.9642, 1.0000, 0.8251)
//
// Arguments of the function may or may not be comma-separated.
//
// The following units are supported for the arguments:
//   - Percents (%)
//   - Radians (rad)
//   - Turns (turn)
//   - Degrees (deg)
//
// Additionally, none is interpreted as zero:
//
//	color.ParseFunc("lch(29.2345% 44.2 none)") // &color.Lch{0.29345, 44.2, 0}
//
// Note that ParseFunc does not implement CSS color functions completely. For
// example, there is no support for relative values; i.e. the following call:
//
//	color.ParseFunc("hsl(from green h s l / 0.5)")
//
// results in error.
func ParseFunc(fnstring string) (Color, error) {
	tokens := strings.FieldsFunc(fnstring, func(r rune) bool {
		return r == ' ' || r == ',' || r == '(' || r == ')'
	})

	if len(tokens) < 1 {
		return nil, fmt.Errorf("failed to tokenize function string")
	}

	var (
		// Name of the function, e.g. rgb, hsl, xyz.
		funcName = tokens[0]

		// Argument slice (arguments stored as strings).
		args = tokens[1:]

		// Expected length of the argument slice.
		//
		// All color functions require at least 3 arguments, with only CMYK
		// requiring more than 3 arguments
		expectedLength = 3
	)

	if funcName == "cmyk" {
		expectedLength++
	}

	if err := checkArgsLen(funcName, expectedLength, args); err != nil {
		return nil, err
	}

	parsedArgs := make([]float64, expectedLength)

	for index, arg := range args {
		parsed, err := parseArg(arg)
		if err != nil {
			return nil, err
		}
		parsedArgs[index] = parsed
	}

	switch funcName {
	case "cmyk":
		return NewCMYK(
			percents.FromFloat(parsedArgs[0]),
			percents.FromFloat(parsedArgs[1]),
			percents.FromFloat(parsedArgs[2]),
			percents.FromFloat(parsedArgs[3]),
		), nil
	case "hsl":
		return NewHSL(
			int(math.Round(parsedArgs[0])),
			percents.FromFloat(parsedArgs[1]),
			percents.FromFloat(parsedArgs[2]),
		), nil
	case "hsv":
		return NewHSV(
			int(math.Round(parsedArgs[0])),
			percents.FromFloat(parsedArgs[1]),
			percents.FromFloat(parsedArgs[2]),
		), nil
	case "lab":
		return NewLab(
			parsedArgs[0],
			parsedArgs[1],
			parsedArgs[2],
		), nil
	case "lch":
		return NewLch(
			parsedArgs[0],
			parsedArgs[1],
			int(math.Round(parsedArgs[2])),
		), nil
	case "rgb":
		return NewRGB(
			uint8(math.Round(parsedArgs[0])),
			uint8(math.Round(parsedArgs[1])),
			uint8(math.Round(parsedArgs[2])),
		), nil
	case "xyz":
		return NewXYZ(
			parsedArgs[0],
			parsedArgs[1],
			parsedArgs[2],
		), nil
	}

	return nil, fmt.Errorf("unknown color function %s", funcName)
}

// parseArg parses the argument.
//
// It supports the following units:
//   - Percents (%)
//   - Radians (rad)
//   - Turns (turn)
//   - Degrees (deg)
//
// Additionally, none is interpreted as zero.
func parseArg(arg string) (float64, error) {
	// None.
	if arg == "none" {
		return 0, nil
	}

	// Percents.
	// In this case the parsed value should be divided by 100.
	if arg[len(arg)-1] == '%' {
		norm := arg[:len(arg)-1]

		v, err := strconv.ParseFloat(norm, 64)
		if err != nil {
			return 0, err
		}

		return v / 100, nil
	}

	// Radians.
	// In this case the parsed value should be converted to degrees.
	if strings.HasSuffix(arg, "rad") {
		norm := arg[:len(arg)-3]

		v, err := strconv.ParseFloat(norm, 64)
		if err != nil {
			return 0, err
		}

		return float64(degrees.FromRadians(v)), nil
	}

	// Turns.
	// In this case the parsed value should be converted to degrees
	if strings.HasSuffix(arg, "turn") {
		norm := arg[:len(arg)-4]

		v, err := strconv.ParseFloat(norm, 64)
		if err != nil {
			return 0, err
		}

		return float64(degrees.Normalize(degrees.FromTurn(v))), nil
	}

	// Degrees.
	if strings.HasSuffix(arg, "deg") {
		norm := arg[:len(arg)-3]
		return strconv.ParseFloat(norm, 64)
	}

	// Other cases.
	return strconv.ParseFloat(arg, 64)
}

// checkArgsLen compares actual length of the slice array with the expected and
// returns formatted error if they don't match.
func checkArgsLen(funcName string, expected int, args []string) error {
	got := len(args)

	if got < expected {
		return fmt.Errorf("not enough arguments for %s: expected %d, got %d", funcName, expected, got)
	}

	if got > expected {
		return fmt.Errorf("too many arguments for %s: expected %d, got %d", funcName, expected, got)
	}

	return nil
}
