package main

import (
	"math/rand"
	"time"

	"github.com/ntsvetko/gobucks/models"
	"gopkg.in/mgo.v2"
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

// returns true for win, false for loss, error for failed operation
func Gamble(name string, bet int, userColl *mgo.Collection) (bool, int, error) {
	outcome := trueOrFalse()
	transaction, currAmt, err := models.AddTransaction(name, bet, outcome, userColl)
	if err != nil {
		return false, currAmt, err
	}

	if transaction == false {
		return false, currAmt, nil
	}

	return true, currAmt, nil
}
