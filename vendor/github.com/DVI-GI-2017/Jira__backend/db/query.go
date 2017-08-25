package db

import "gopkg.in/mgo.v2"

// Interface for generic queries.
type Query interface {
	// Returns all query entries in result object.
	// NOTE: Result MUST be POINTER to slice of objects with specified type.
	All(result interface{}) error

	// Return first object in query
	// NOTE: Result MUST be POINTER to object with specified type.
	One(result interface{}) error

	// Count all objects in query
	Count() (count int, err error)

	// Returns query with fields specified in selector
	Select(selector interface{}) Query

	// Returns true if query is empty
	IsEmpty() (empty bool, err error)
}

// Wrapper around *mgo.Query
type MongoQuery struct {
	*mgo.Query
}

// Returns true if query is empty
func (q MongoQuery) IsEmpty() (empty bool, err error) {
	count, err := q.Count()
	return count == 0, err
}

// Returns query selected from initial by selector.
func (q MongoQuery) Select(selector interface{}) Query {
	return MongoQuery{q.Query.Select(selector)}
}
