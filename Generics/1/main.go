package main

func getLast[T any](s []T) T {
	var lastValue T
	if	0 == len(s){ 
		return lastValue
	}

	return s[len(s)-1]
}
