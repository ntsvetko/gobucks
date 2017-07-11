package main

import (
	"math/rand"
	"time"

	"github.com/ntsvetko/gobucks/models"
	"gopkg.in/mgo.v2"
)

/*
* psuedo-random gambling method
 */
func trueOrFalse() bool {
	seed := rand.NewSource(time.Now().UnixNano())
	seededRand := rand.New(seed)
	var i = seededRand.Intn(100)
	return (i >= 50)
}

/*Gamble takes a user name, bet amount, mongo session, database name and collection name
* If the user does not exist, one will be created with a starting fund of 100
*
* returns true for win, false for loss, current user balance and an error for failed operation
 */
func Gamble(name string, bet int, session *mgo.Session, dbName string, collName string) (bool, int, error) {
	outcome := trueOrFalse()

	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	userColl := sessionCopy.DB(dbName).C(collName)
	transaction, currAmt, err := models.AddTransaction(name, bet, outcome, userColl)
	if err != nil {
		return false, currAmt, err
	}

	if transaction == false {
		return false, currAmt, nil
	}

	return true, currAmt, nil
}

/*Balance returns the current balance of the specified user
 */
func Balance(name string, session *mgo.Session, dbName string, collName string) int {
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	userColl := sessionCopy.DB(dbName).C(collName)

	user := models.FindOrCreateUser(name, userColl)

	return user.CurrAmount
}
