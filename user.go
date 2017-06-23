package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name               string
	TransactionHistory []Transaction
	CurrAmount         int
}

// creates and adds user to db, does NOT check if user already exists
func createUser(name string, startingBalance int, userColl *mgo.Collection) (*User, error) {
	user := User{name, []Transaction{}, startingBalance}
	err := userColl.Insert(&User{name, []Transaction{}, startingBalance})

	return &user, err
}

func findUser(name string, userColl *mgo.Collection) (*User, error) {
	user := User{}
	err := userColl.Find(bson.M{"Name": name}).One(&user)

	return &user, err
}

// if user exists, returns that user, otherwise creates and returns new user
func findOrCreateUser(name string, userColl *mgo.Collection) *User {
	user, err := findUser(name, userColl)

	/*if err != nil {
		log.Fatal(err)
	}*/

	if user != nil {
		return user
	}

	user, err = createUser(name, 100, userColl)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
