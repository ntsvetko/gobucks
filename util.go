package main

import mgo "gopkg.in/mgo.v2"

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

func seedUsers(dbConf dbConfig, session *mgo.Session) ([]string, *mgo.Collection) {
	userColl := session.DB(dbConf.dbName).C(dbConf.collName)
	names := []string{"jae", "marcus", "isaiah", "jonas", "kelly", "jordan", "jimmy", "bob", "roger", "bill"}

	var users = [10]*User{}

	for i := 0; i < len(names); i++ {
		users[i] = findOrCreateUser(names[i], userColl)
	}

	return names, userColl
}
