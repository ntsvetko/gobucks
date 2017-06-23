package main

import (
	"testing"
)

func TestUpdateTransaction(test *testing.T) {
	dbConf, session := initResetDB()
	userColl := seedUsers(dbConf, session)

	dbSize, cErr := userColl.Count()

	if cErr != nil {
		test.Fatalf("Database error in call to collections.Count(), error should be nil, but was %v", cErr)
	}
	if dbSize != 10 {
		test.Fatalf("Database should have 10 element in it, but has %v", dbSize)
	}
}
