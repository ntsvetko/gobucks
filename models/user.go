package models

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

//returns user if found, err if user doesnt exist
func findUser(name string, userColl *mgo.Collection) (*User, error) {
	user := User{}
	err := userColl.Find(bson.M{"name": name}).One(&user)

	return &user, err
}

/*FindOrCreateUser takes a username and returns the user if found in DB, creates and inserts new one otherwise
* returns user created/found
 */
func FindOrCreateUser(name string, userColl *mgo.Collection) *User {
	user, err := findUser(name, userColl)

	if err == nil {
		return user
	}

	user, err = createUser(name, 100, userColl)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
