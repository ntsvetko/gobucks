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

	// add winning transaction to user with empty history

	user := User{}
	err := userColl.Find(bson.M{"name": names[0]}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}
	beforeAmt, betAmt := 100, 15
	outcome := true

	r, _ := AddTransaction(user.Name, betAmt, outcome, userColl)

	if r != true {
		test.Fatalf("AddTransaction should return true for successful operations")
	}

	//get updated user
	err = userColl.Find(bson.M{"name": names[0]}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}

	if len(user.TransactionHistory) != 1 {
		test.Fatalf("length of transactionhistory should be 1 but was %v", err)
	}

	trans := user.TransactionHistory[0]

	if trans.AmountBefore != beforeAmt {
		test.Fatalf("trans.amountBefore should be "+strconv.Itoa(beforeAmt)+" but was %v", trans.AmountBefore)
	}

	if trans.AmountGambled != betAmt {
		test.Fatalf("trans.amountGambled should be "+strconv.Itoa(betAmt)+" but was %v", trans.AmountGambled)
	}

	if trans.Result != outcome {
		test.Fatalf("trans.result should be "+strconv.FormatBool(outcome)+" but was %v", trans.Result)
	}

	if trans.AmountAfter != (beforeAmt + betAmt) {
		test.Fatalf("trans.amuntAfter should be "+strconv.Itoa(beforeAmt+betAmt)+" but was %v", trans.AmountAfter)
	}

	if user.CurrAmount != (beforeAmt + betAmt) {
		test.Fatalf("user.CurrAmount should be "+strconv.Itoa(beforeAmt+betAmt)+" but was %v", user.CurrAmount)
	}

	// add winning transaction to non-empty history

	beforeAmt, betAmt = 115, 30
	outcome = true

	AddTransaction(user.Name, betAmt, outcome, userColl)

	err = userColl.Find(bson.M{"name": names[0]}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}

	if len(user.TransactionHistory) != 2 {
		test.Fatalf("length of transactionhistory should be 1 but was %v", err)
	}

	trans = user.TransactionHistory[1]

	if trans.AmountBefore != beforeAmt {
		test.Fatalf("trans.amountBefore should be "+strconv.Itoa(beforeAmt)+" but was %v", trans.AmountBefore)
	}

	if trans.AmountGambled != betAmt {
		test.Fatalf("trans.amountGambled should be "+strconv.Itoa(betAmt)+" but was %v", trans.AmountGambled)
	}

	if trans.Result != outcome {
		test.Fatalf("trans.result should be "+strconv.FormatBool(outcome)+" but was %v", trans.Result)
	}

	if trans.AmountAfter != (beforeAmt + betAmt) {
		test.Fatalf("trans.amuntAfter should be "+strconv.Itoa(beforeAmt+betAmt)+" but was %v", trans.AmountAfter)
	}

	if user.CurrAmount != (beforeAmt + betAmt) {
		test.Fatalf("user.CurrAmount should be "+strconv.Itoa(beforeAmt+betAmt)+" but was %v", user.CurrAmount)
	}

	// add losing transaction to non-empty history

	beforeAmt, betAmt = 145, 35
	outcome = false

	AddTransaction(user.Name, betAmt, outcome, userColl)

	err = userColl.Find(bson.M{"name": names[0]}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}

	if len(user.TransactionHistory) != 3 {
		test.Fatalf("length of transactionhistory should be 1 but was %v", err)
	}

	trans = user.TransactionHistory[2]

	if trans.AmountBefore != beforeAmt {
		test.Fatalf("trans.amountBefore should be "+strconv.Itoa(beforeAmt)+" but was %v", trans.AmountBefore)
	}

	if trans.AmountGambled != betAmt {
		test.Fatalf("trans.amountGambled should be "+strconv.Itoa(betAmt)+" but was %v", trans.AmountGambled)
	}

	if trans.Result != outcome {
		test.Fatalf("trans.result should be "+strconv.FormatBool(outcome)+" but was %v", trans.Result)
	}

	if trans.AmountAfter != (beforeAmt - betAmt) {
		test.Fatalf("trans.amuntAfter should be "+strconv.Itoa(beforeAmt+betAmt)+" but was %v", trans.AmountAfter)
	}

	if user.CurrAmount != (beforeAmt - betAmt) {
		test.Fatalf("user.CurrAmount should be "+strconv.Itoa(beforeAmt+betAmt)+" but was %v", user.CurrAmount)
	}

	// add losing transaction to empty history

	beforeAmt, betAmt = 100, 35
	outcome = false

	//change user to one with empty transactionhistory
	err = userColl.Find(bson.M{"name": names[1]}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}

	AddTransaction(user.Name, betAmt, outcome, userColl)

	err = userColl.Find(bson.M{"name": names[1]}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}

	if len(user.TransactionHistory) != 1 {
		test.Fatalf("length of transactionhistory should be 1 but was %v", err)
	}

	trans = user.TransactionHistory[0]

	if trans.AmountBefore != beforeAmt {
		test.Fatalf("trans.amountBefore should be "+strconv.Itoa(beforeAmt)+" but was %v", trans.AmountBefore)
	}

	if trans.AmountGambled != betAmt {
		test.Fatalf("trans.amountGambled should be "+strconv.Itoa(betAmt)+" but was %v", trans.AmountGambled)
	}

	if trans.Result != outcome {
		test.Fatalf("trans.result should be "+strconv.FormatBool(outcome)+" but was %v", trans.Result)
	}

	if trans.AmountAfter != (beforeAmt - betAmt) {
		test.Fatalf("trans.amuntAfter should be "+strconv.Itoa(beforeAmt+betAmt)+" but was %v", trans.AmountAfter)
	}

	if user.CurrAmount != (beforeAmt - betAmt) {
		test.Fatalf("user.CurrAmount should be "+strconv.Itoa(beforeAmt+betAmt)+" but was %v", user.CurrAmount)
	}

	// call AddTransaction with non-existing user (should call FindOrCreate user and return successfully)

	r, _ = AddTransaction("randomNameNotInDBm", 5, false, userColl)

	if r == false {
		test.Fatal("AddTransaction should create a user and return true for calls with users that don't yet exist")
	}

	err = userColl.Find(bson.M{"name": "randomNameNotInDBm"}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}

	if len(user.TransactionHistory) != 1 {
		test.Fatalf("length of transactionhistory should be 1 but was %v", err)
	}

	trans = user.TransactionHistory[0]

	if trans.AmountBefore != 100 {
		test.Fatalf("trans.amountBefore should be "+strconv.Itoa(100)+" but was %v", trans.AmountBefore)
	}

	if trans.AmountGambled != 5 {
		test.Fatalf("trans.amountGambled should be "+strconv.Itoa(5)+" but was %v", trans.AmountGambled)
	}

	if trans.Result != false {
		test.Fatalf("trans.result should be "+strconv.FormatBool(false)+" but was %v", trans.Result)
	}

	if trans.AmountAfter != 95 {
		test.Fatalf("trans.amuntAfter should be "+strconv.Itoa(95)+" but was %v", trans.AmountAfter)
	}

	if user.CurrAmount != 95 {
		test.Fatalf("user.CurrAmount should be "+strconv.Itoa(95)+" but was %v", user.CurrAmount)
	}

	// invalid calls to AddTransaction

	err = userColl.Find(bson.M{"name": names[2]}).One(&user)
	if err != nil {
		test.Fatalf("database error in TestAddTransaction, err should be nil but was %v", err)
	}

	beforeAmt, betAmt = user.CurrAmount, (user.CurrAmount + 10)
	outcome = false

	r, _ = AddTransaction(user.Name, betAmt, outcome, userColl)

	if r == true {
		test.Fatal("AddTransaction should return false for invalid transactions")
	}

}
