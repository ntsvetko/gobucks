package main

import (
	"math/rand"
	"time"
)

func trueOrFalse() bool {
	seed := rand.NewSource(time.Now().UnixNano())
	seededRand := rand.New(seed)
	var i = seededRand.Intn(100)
	if i < 50 {
		return true
	}

	return false
}
