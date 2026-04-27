
package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	if "pro"==plan{
		return messages[:],nil
	}
	if "free"==plan {
		return messages[0:2],nil
	}
	return nil,errors.New("unsupported plan")
}
