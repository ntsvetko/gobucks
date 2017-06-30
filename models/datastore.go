package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

// methods that call ConnectToMongo needs to close session when done
func ConnectToMongo(uri string) *mgo.Session {
	session, err := mgo.Dial(uri)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongo")

	return session
}

func GetColl(session *mgo.Session, dbName string, collName string) *mgo.Collection {
	coll := session.DB(dbName).C(collName)
	return coll
}
