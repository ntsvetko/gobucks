package models

import (
	"testing"

	"strconv"

	"gopkg.in/mgo.v2/bson"
)

func TestAddTransaction(test *testing.T) {
	dbConf, session := initResetDB()
	names, userColl := seedUsers(dbConf, session)

	dbSize, cErr := userColl.Count()

	if cErr != nil {
		test.Fatalf("Database error in call to collections.Count(), error should be nil, but was %v", cErr)
	}
	if dbSize != 10 {
		test.Fatalf("Database should have 10 element in it, but has %v", dbSize)
	}

	// add transaction to user with empty history

	user := User{}
	err := userColl.Find(bson.M{"name": names[0]}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}
	beforeAmt, betAmt := user.CurrAmount, 15
	outcome := true

	addTransaction(user.Name, beforeAmt, betAmt, outcome)

	err = userColl.Find(bson.M{"name": names[0]}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}

	if len(user.TransactionHistory) != 1 {
		test.Fatalf("length of transactionhistory should be 1 but was %v", err)
	}

	trans := user.TransactionHistory[0]

	if trans.amountBefore != beforeAmt {
		test.Fatalf("trans.amountBefore should be "+strconv.Itoa(beforeAmt)+" but was %v", trans.amountBefore)
	}

	if trans.amountGambled != betAmt {
		test.Fatalf("trans.amountGambled should be "+strconv.Itoa(betAmt)+" but was %v", trans.amountGambled)
	}

	if trans.result != outcome {
		test.Fatalf("trans.result should be "+strconv.FormatBool(outcome)+" but was %v", trans.result)
	}

	if trans.amountAfter != (beforeAmt + betAmt) {
		test.Fatalf("trans.amuntAfter should be "+strconv.Itoa(beforeAmt+betAmt)+" but was %v", trans.amountAfter)
	}

}
