package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTrueOrFalse(test *testing.T) {

	var trueFalseCount [2]int
	for i := 0; i < 100; i++ {
		if trueOrFalse() == true {
			trueFalseCount[1]++
		} else {
			trueFalseCount[0]++
		}
	}

	fmt.Println("Number of false results: " + strconv.Itoa(trueFalseCount[0]))
	fmt.Println("Number of true results: " + strconv.Itoa(trueFalseCount[1]))
}
