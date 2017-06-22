package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func connectToMongo(uri string) *mgo.Session {
	session, err := mgo.Dial(uri)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongo")

	return session
}

func getColl(session *mgo.Session, dbName string, collName string) *mgo.Collection {
	coll := session.DB(dbName).C(collName)
	return coll
}
