package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type Transaction struct {
	AmountBefore  int
	AmountGambled int
	Result        bool
	AmountAfter   int
}

/*AddTransaction takes a any string as a username, any bet amount greater than 0
* and less than the user's balance
*
* if the user specified does not already exist, one will be created with 100 gobucks
*
* returns true for wins, false for loss, current user balance, and an error for failed transactions
*
 */
func AddTransaction(name string, betAmt int, outcome bool, userColl *mgo.Collection) (bool, int, error) {
	user := FindOrCreateUser(name, userColl)

	currAmt := user.CurrAmount
	var newAmt int
	if betAmt < 0 || betAmt > currAmt {
		return false, currAmt, errors.New(name + " cannot bet " + strconv.Itoa(betAmt) + ". Balance is " + strconv.Itoa(currAmt) + ".") // bad arg
	}

	if outcome == true {
		newAmt = (currAmt + betAmt)
	} else {
		newAmt = (currAmt - betAmt)
	}

	if newAmt < 0 {
		return false, currAmt, errors.New("bad args") // invalid args
	}

	trans := Transaction{currAmt, betAmt, outcome, newAmt}

	err := userColl.Update(bson.M{"name": name}, bson.M{"$push": bson.M{"transactionhistory": trans}})

	if err != nil {
		return false, currAmt, err
	}

	err = userColl.Update(bson.M{"name": name}, bson.M{"$set": bson.M{"curramount": newAmt}})

	if err != nil {
		return false, currAmt, err
	}

	return outcome, newAmt, nil
}
