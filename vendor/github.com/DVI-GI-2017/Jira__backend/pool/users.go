package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/users"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	resolvers["User"] = usersResolver
}

const (
	UserCreate      = Action("UserCreate")
	UserExists      = Action("UserExists")
	UserAuthorize   = Action("UserAuthorize")
	UserFindById    = Action("UserFindById")
	UserFindByEmail = Action("UserFindByEmail")
	UsersAll        = Action("UsersAll")
	UserAllProjects = Action("UserAllProjects")
)

func usersResolver(action Action) (service ServiceFunc) {
	switch action {

	case UserCreate:
		service = func(source db.DataSource, user interface{}) (result interface{}, err error) {
			return users.CreateUser(source, user.(models.User))
		}
		return

	case UserExists:
		service = func(source db.DataSource, credentials interface{}) (result interface{}, err error) {
			return users.CheckUserExists(source, credentials.(models.User))
		}
		return

	case UserAuthorize:
		service = func(source db.DataSource, credentials interface{}) (interface{}, error) {
			return users.CheckUserCredentials(source, credentials.(models.User))
		}
		return

	case UsersAll:
		service = func(source db.DataSource, _ interface{}) (result interface{}, err error) {
			return users.AllUsers(source)
		}
		return

	case UserFindById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return users.FindUserById(source, id.(bson.ObjectId))
		}
		return

	case UserFindByEmail:
		service = func(source db.DataSource, email interface{}) (interface{}, error) {
			return users.FindUserByEmail(source, email.(models.Email))
		}
		return

	case UserAllProjects:
		service = func(source db.DataSource, id interface{}) (result interface{}, err error) {
			return users.AllUsersProject(source, id.(models.RequiredId))
		}
		return

	default:
		log.Panicf("can not find resolver with action: %v, in users resolvers", action)
		return
	}
}
