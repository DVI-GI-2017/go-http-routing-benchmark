package pool

import (
	"strings"

	"github.com/DVI-GI-2017/Jira__backend/db"
)

type Action string

// Returns true if action has prefix "prefix"
func (a Action) HasPrefix(prefix string) bool {
	return strings.HasPrefix(string(a), prefix)
}

// Function to be passed to workers pool
type ServiceFunc func(source db.DataSource, data interface{}) (result interface{}, err error)

// Function that takes action and returns service function
type ResolverFunc func(action Action) (service ServiceFunc, err error)

// Resolvers will be initialized from files with resolvers in this package
var resolvers = make(map[string]ResolverFunc, 0)
