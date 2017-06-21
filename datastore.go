package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func connectToMongo(url string) *mgo.Session {
	session, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Println("Connected to mongo")

	return session
}
