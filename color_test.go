package color_test

import (
	"fmt"

	"github.com/shelepuginivan/color"
)

func Example() {
	// Parse color from CSS functions.
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
