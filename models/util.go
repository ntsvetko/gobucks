package models

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

func InitResetDB() (dbConfig, *mgo.Session) {
	dbConf := getDbConfig()
	session := ConnectToMongo(dbConf.uri)

	session.DB(dbConf.dbName).DropDatabase()
	return dbConf, session
}

func SeedUsers(dbConf dbConfig, collName string, session *mgo.Session) ([]string, *mgo.Collection) {
	userColl := session.DB(dbConf.dbName).C(collName)
	names := []string{"jae", "marcus", "isaiah", "jonas", "kelly", "jordan", "jimmy", "bob", "roger", "bill"}

	var users = [10]*User{}

	for i := 0; i < len(names); i++ {
		users[i] = FindOrCreateUser(names[i], userColl)
	}

	return names, userColl
}
