package db

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/configs"
)

// Initialize global defaultDB instance
func InitDB(mongo *configs.Mongo) {
	log.Println("Connecting to local mongo server....")

	session, err := NewMongoSession(mongo.URL())
	if err != nil {
		log.Panicf("can not connect to mongo server: %v", err)
	}

	defaultDB = session.Source(mongo.DB)
}

// Default defaultDB
var defaultDB DataSource

// Returns current data source with new session
func Copy() DataSource {
	return defaultDB.Copy()
}

// Returns current data source.
func Get() DataSource {
	return defaultDB
}
