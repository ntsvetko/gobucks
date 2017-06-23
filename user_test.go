package main

import (
	"testing"

	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type dbConfig struct {
	uri      string
	dbName   string
	collName string
}

func getDbConfig() dbConfig {
	dbConf := dbConfig{"localhost:27017", "test", "users"}

	return dbConf
}

func initResetDB() (dbConfig, *mgo.Session) {
	dbConf := getDbConfig()
	session := connectToMongo(dbConf.uri)

	session.DB(dbConf.dbName).DropDatabase()
	return dbConf, session
}

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
		test.Fatalf("createUser should return a user, but instead returned %v", err)
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

// func TestFindUser(test *testing.T) {
// 	dbConf, session := initResetDB()
// 	userColl := session.DB(dbConf.dbName).C(dbConf.collName)
// 	userName := "Avery"

// 	user, err := createUser(username, 100, userColl)
// }
