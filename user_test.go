package main

import (
	"testing"

	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func TestCreateUser(test *testing.T) {
	dbConf, session := initResetDB()
	userColl := session.DB(dbConf.dbName).C(dbConf.collName)
	userName := "Danny"
	currAmount := 100
	emptyTransactionHistory := []Transaction{}

	user, err := createUser(userName, currAmount, userColl)

	if err != nil {
		test.Fatalf("database error in createUser, error should be nil but was: %v", err)
	}

	fmt.Println("@createUser: user added: " + user.Name)

	if user == nil {
		test.Fatal("createUser should return a user, but instead returned nil", err)
	}

	if (*user).Name != userName {
		test.Fatalf("user returned was expected to be "+userName+" but was %v", (*user).Name)
	}

	result := User{}
	resErr := userColl.Find(bson.M{"name": userName, "curramount": currAmount, "transactionhistory": emptyTransactionHistory}).One(&result)

	if resErr != nil {
		test.Fatalf("database error in createUser, error should be nil but was: %v", resErr)
	}

	if result.Name != userName {
		test.Fatalf("user returned was expected to be "+userName+" but was %v", result.Name)
	}
}

func TestFindUser(test *testing.T) {
	dbConf, session := initResetDB()
	userColl := session.DB(dbConf.dbName).C(dbConf.collName)
	userName := "Avery"

	createUser(userName, 100, userColl)

	user, err := findUser(userName, userColl)

	if err != nil {
		test.Fatalf("Database error in findUser, error should be nil but was %v", err)
	}

	if user == nil {
		test.Fatal("findUser should return a user, but isntead returned nil", err)
	}

	if (*user).Name != userName {
		test.Fatalf("user returned was supposed to be "+userName+" but was %v", (*user).Name)
	}

}

func TestFindOrCreateUser(test *testing.T) {
	dbConf, session := initResetDB()
	userColl := session.DB(dbConf.dbName).C(dbConf.collName)
	names := []string{"jae", "marcus", "isaiah", "jonas", "kelly", "jordan", "jimmy", "bob", "roger", "bill"}

	// first try with empty collection, should create user
	user := findOrCreateUser(names[0], userColl)

	if user == nil {
		test.Fatal("findOrCreateUser should return a user, but instead returned nil", user)
	}

	if (*user).Name != names[0] {
		test.Fatalf("findOrCreateUser should return user with name: "+(*user).Name+" but instead returned user with name %v", (*user).Name)
	}

	res1, err := findUser(names[0], userColl)

	if err != nil {
		test.Fatalf("Database error in findUser, should be nil but was %v", err)
	}

	if (*res1).Name != names[0] {
		test.Fatalf("user returned was supposed to be "+names[0]+" but was %v", (*res1).Name)
	}

	// calling the method with the same user should return the user without creating a new one

	user = findOrCreateUser(names[0], userColl)

	dbSize, cErr := userColl.Count()

	if cErr != nil {
		test.Fatalf("Database error in call to collections.Count(), error should be nil, but was %v", cErr)
	}
	if dbSize != 1 {
		test.Fatalf("Database should have 1 element in it, but has %v", dbSize)
	}
	res2, err := findUser(names[0], userColl)

	if (*res2).Name != names[0] {
		test.Fatalf("user returned was supposed to be "+names[0]+" but was %v", (*res2).Name)
	}

	// calling this method now with a different user should result in 2 users in the collections

	user = findOrCreateUser(names[1], userColl)

	if user == nil {
		test.Fatal("findOrCreateUser should return a user, but instead returned nil", user)
	}

	if (*user).Name != names[1] {
		test.Fatalf("findOrCreateUser should return user with name: "+(*user).Name+" but instead returned user with name %v", (*user).Name)
	}

	res3, err := findUser(names[1], userColl)

	if err != nil {
		test.Fatalf("Database error in findUser, should be nil but was %v", err)
	}

	if (*res3).Name != names[1] {
		test.Fatalf("user returned was supposed to be "+names[0]+" but was %v", (*res3).Name)
	}

	// add all users from the names array, should get size of array as count

	var users = [10]*User{}

	for i := 0; i < len(names); i++ {
		users[i] = findOrCreateUser(names[i], userColl)
	}

	dbSize, cErr = userColl.Count()

	if cErr != nil {
		test.Fatalf("Database error in call to collections.Count(), error should be nil, but was %v", cErr)
	}
	if dbSize != 10 {
		test.Fatalf("Database should have 10 element in it, but has %v", dbSize)
	}

}
