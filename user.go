package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type User struct {
	name               string
	transactionHistory []Transaction
	currAmount         int
}

// creates and adds user to db, does NOT check if user already exists
func createUser(name string, startingBalance int, userColl *mgo.Collection) (*User, error) {
	user := &User{name, []Transaction{}, startingBalance}
	fmt.Println(user)
	err := userColl.Insert(user)
	return user, err
}

func findUser(name string, userColl *mgo.Collection) (*User, error) {
	user := User{}
	err := userColl.Find(bson.M{"name": name}).One(&user)

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
