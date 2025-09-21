# `color`

Package color provides a robust implementation for color math as well as color manipulation methods.

Here are some of its capabilities at a glance:

```go
package main

import (
	"fmt"

	"github.com/shelepuginivan/color"
)

func main() {
	// Parse colors from CSS functions.
	scarlet := color.Must(color.Parse("hsl(8deg, 100%, 50%)"))

	// Generate shades, tints, and tones.
	for _, tone := range color.Tones(scarlet, 5) {
		// Color wheel functions.
		c1, c2 := color.SplitComplementary(tone)

		fmt.Println(tone.Hex(), c1.Hex(), c2.Hex())
	}

	mint := color.NewRGB(209, 227, 217)

	// Mix colors in different colorspaces.
	peach := color.MixLab(scarlet, mint)

	// Convert colors to different colorspaces.
	fmt.Println(peach.XYZWithWhitepoint(color.D55))
}
```

## Features

- Consistent chainable API
- Parsing from CSS-like color functions
- Color convertion and mixing
- Lightness and contrast calculation
- Color wheel functions

## Installation

```shell
go get github.com/shelepuginivan/color
```

## Documentation

Is available on [pkg.go.dev](https://pkg.go.dev/github.com/shelepuginivan/color).
