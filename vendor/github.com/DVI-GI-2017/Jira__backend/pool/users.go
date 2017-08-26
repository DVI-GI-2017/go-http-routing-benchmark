package pool

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services"
)

func init() {
	resolvers["User"] = usersResolver
}

const (
	UserCreate      = Action("UserCreate")
	UserExists      = Action("UserExists")
	UserAuthorized  = Action("UserAuthorized")
	UserFindById    = Action("UserFindById")
	UserFindByEmail = Action("UserFindByEmail")
	UsersAll        = Action("UsersAll")
	UserAllProjects = Action("UserAllProjects")
)

func usersResolver(action Action) (service ServiceFunc, err error) {
	switch action {

	case UserCreate:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			user, err := models.SafeCastToUser(data)
			if err != nil {
				return models.User{}, err
			}
			return services.CreateUser(source, user)
		}
		return

	case UserExists:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			user, err := models.SafeCastToUser(data)
			if err != nil {
				return false, err
			}
			return services.CheckUserExists(source, user)
		}
		return

	case UserAuthorized:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			user, err := models.SafeCastToUser(data)
			if err != nil {
				return false, err
			}
			return services.AuthorizeUser(source, user)
		}
		return

	case UsersAll:
		service = func(source db.DataSource, _ interface{}) (result interface{}, err error) {
			return services.AllUsers(source)
		}
		return

	case UserFindById:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return models.User{}, err
			}
			return services.FindUserById(source, id)
		}
		return

	case UserFindByEmail:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			email, err := models.SafeCastToEmail(data)
			if err != nil {
				return models.User{}, err
			}
			return services.FindUserByEmail(source, email)
		}
		return

	case UserAllProjects:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return models.ProjectsList{}, err
			}
			return services.AllUserProjects(source, id)
		}
		return
	}
	return nil, fmt.Errorf("can not find resolver with action: %v, in users resolvers", action)
}
