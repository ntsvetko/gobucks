package main

import (
	"testing"

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
	userName := "vlad"

	user, err := createUser(userName, 100, userColl)

	if err != nil {
		test.Fatalf("database error in createUser, error should be nil but was: %v", err)
	}

	if user == nil {
		test.Fatalf("createUser should return a user, but instead returned %v", err)
	}

	if (*user).name != userName {
		test.Fatalf("user returned was expected to be "+userName+" but was %v", (*user).name)
	}

	result := User{}
	resErr := userColl.Find(bson.M{"name": userName}).One(&result)

	if resErr != nil {
		test.Fatalf("database error in createUser, error should be nil but was: %v", resErr)
	}

	if result.name != userName {
		test.Fatalf("user returned was expected to be "+userName+" but was %v", result.name)
	}
}
