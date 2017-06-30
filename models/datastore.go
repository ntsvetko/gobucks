package models

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

/*ConnectToMongo returns a mongodb session
* only call this once, use session.Copy() after
* methods that call ConnectToMongo needs to close session when done
 */
func ConnectToMongo(uri string) *mgo.Session {
	session, err := mgo.Dial(uri)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongo")

	return session
}
