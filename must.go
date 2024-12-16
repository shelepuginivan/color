package color

// Must is a helper that wraps a call to a function returning (Color, error)
// and panics if the error is not nil. It is intended for use in variable
// initializations such as
//
//	c := color.Must(color.ParseHex("#ffffff"))
func Must(c Color, err error) Color {
	if err != nil {
		panic(err)
	}
	return c
}
