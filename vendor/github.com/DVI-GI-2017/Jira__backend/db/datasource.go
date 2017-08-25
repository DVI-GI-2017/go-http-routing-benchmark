package db

import "gopkg.in/mgo.v2"

// Interface for generic data source (e.g. database).
type DataSource interface {
	// Returns collection by name
	C(name string) Collection

	// Returns copy of data source (may be copy of session as well)
	Copy() DataSource

	// Closes data source (it will be runtime error to use it after close)
	Close()
}

// Override Source method of mgo.Session to return wrapper around *mgo.DataSource.
func (s MongoSession) Source(name string) DataSource {
	return &MongoDatabase{Database: s.Session.DB(name)}
}

// Wrapper around *mgo.DataSource.
type MongoDatabase struct {
	*mgo.Database
}

// Override C method of mgo.DataSource to return wrapper around *mgo.Collection
func (d MongoDatabase) C(name string) Collection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

// Returns database associated with copied session
func (d MongoDatabase) Copy() DataSource {
	return MongoDatabase{d.With(d.Session.Copy())}
}

// Closes current session with mongo db
func (d MongoDatabase) Close() {
	d.Database.Session.Close()
}
