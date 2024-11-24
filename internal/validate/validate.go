package validate

import "fmt"

// IsPercent validates value in percents.
func IsPercent(p int, argname string) error {
	if p < 0 || p > 100 {
		return fmt.Errorf("%s must be a valid value in percents (integer in range [0, 100]), got %d", argname, p)
	}
	return nil
}

// IsDegree validates value in degrees.
func IsDegree(d int, argname string) error {
	if d < 0 || d > 360 {
		return fmt.Errorf("%s must be a valid value in degrees (integer in range [0, 360]), got %d", argname, d)
	}
	return nil
}

// All checks all errors provided as arguments and returns the first non-nil
// error value. If there are no errors, nil is returned.
func All(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
