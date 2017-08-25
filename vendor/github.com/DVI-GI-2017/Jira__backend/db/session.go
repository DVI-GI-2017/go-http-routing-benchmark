package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// Interface for generic session with data source.
type Session interface {
	// Returns data source associated with this session with "name"
	Source(name string) DataSource

	// Closes session with data source other methods should panic
	// if you are trying to use closed session
	Close()
}

// Wrapper around *mgo.Session.
type MongoSession struct {
	*mgo.Session
}

// Creates new mongo session
func NewMongoSession(mgoURI string) (Session, error) {
	mgoSession, err := mgo.Dial(mgoURI)
	if err != nil {
		return nil, fmt.Errorf("can not open defaultDB session: %v", err)
	}

	mgoSession.SetMode(mgo.Monotonic, true)

	return MongoSession{mgoSession}, nil
}
