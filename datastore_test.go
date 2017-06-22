package main

import (
	mgo "gopkg.in/mgo.v2"
)

func cnnectToMongo() *mgo.Session {
	session := connectToMongo("localhost:27017")

	return session
}
