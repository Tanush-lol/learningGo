
package main

import (
	"errors"
)

func validateStatus(status string) error {
	var length int= len(status)
	if 0==length {
		return errors.New("status cannot be empty")
	}
	if 140<length {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}
