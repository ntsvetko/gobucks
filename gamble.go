package main

import (
	"github.com/ntsvetko/gobucks/models"
	"gopkg.in/mgo.v2"
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

// returns true for success, false for failure
func Gamble(name string, bet int, userColl *mgo.Collection) bool {
	outcome := trueOrFalse()
	transactionSuccess, err := models.AddTransaction(name, bet, outcome, userColl)
	if err != nil {
		return false
	}

	if transactionSuccess == false {
		return false
	}

	return true
}
