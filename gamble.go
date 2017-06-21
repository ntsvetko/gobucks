package main

import "math/rand"

func trueOrFalse() bool {
	var i = rand.Intn(100)
	if i < 50 {
		return true
	}

	return false
}
