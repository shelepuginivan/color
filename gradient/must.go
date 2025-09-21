package gradient

// Must is a helper that wraps a call to a function returning (Gradient, error)
// and panics if the error is not nil. It is intended for use in variable
// initializations such as
//
//	c := gradient.Must(gradient.NewLinear(
//	    gradient.InOklch(gradient.LongerHue),
//	    gradient.WithColorStop(c1, 0),
//	    gradient.WithColorStop(c2, 0),
//	))
func Must(c Gradient, err error) Gradient {
	if err != nil {
		panic(err)
	}
	return c
}

