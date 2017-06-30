package models

import "gopkg.in/mgo.v2/bson"
import "gopkg.in/mgo.v2"

type Transaction struct {
	AmountBefore  int
	AmountGambled int
	Result        bool
	AmountAfter   int
}

// returns true for successful additions to the history, false otherwise
func AddTransaction(name string, betAmt int, outcome bool, userColl *mgo.Collection) (bool, error) {
	if betAmt < 0 {
		return false, nil // bad arg
	}

	user := FindOrCreateUser(name, userColl)

	currAmt := user.CurrAmount
	var newAmt int

	if outcome == true {
		newAmt = (currAmt + betAmt)
	} else {
		newAmt = (currAmt - betAmt)
	}

	if newAmt < 0 {
		return false, nil // invalid args
	}

	trans := Transaction{currAmt, betAmt, outcome, newAmt}

	err := userColl.Update(bson.M{"name": name}, bson.M{"$push": bson.M{"transactionhistory": trans}})

	if err != nil {
		return false, err
	}

	err = userColl.Update(bson.M{"name": name}, bson.M{"$set": bson.M{"curramount": newAmt}})

	if err != nil {
		return false, err
	}

	return true, nil
}
