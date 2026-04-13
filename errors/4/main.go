
package main

import (
	"errors"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		// ?
		return 0,errors.New("no dividing by 0")
//Remember that it's conventional to return the "zero" values of all other return values when you return a non-nil error in Go.
	}
	return x / y, nil
}
