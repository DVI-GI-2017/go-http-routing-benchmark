package db

import "gopkg.in/mgo.v2"

// Interface for generic collection.
type Collection interface {
	// Returns Query with objects with given id
	FindId(id interface{}) Query

	// Returns Query that build atop of given query spec.
	Find(query interface{}) Query

	// Inserts object.
	Insert(value interface{}) error

	// Finds object with "selector" and update it using "update" object.
	Update(selector, update interface{}) error
}

// Wrapper around *mgo.Collection
type MongoCollection struct {
	*mgo.Collection
}

// Return Query with objects with given id
func (c MongoCollection) FindId(id interface{}) Query {
	return MongoQuery{c.Collection.FindId(id)}
}

// Return Query with objects that match given query
func (c MongoCollection) Find(query interface{}) Query {
	return MongoQuery{c.Collection.Find(query)}
}

// Inserts value in collection.
func (c MongoCollection) Insert(value interface{}) error {
	return c.Collection.Insert(value)
}

// Updates collection documents found by selector with "update" document.
func (c MongoCollection) Update(selector, update interface{}) error {
	return c.Collection.Update(selector, update)
}
