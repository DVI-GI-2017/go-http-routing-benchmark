package pool

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services"
)

func init() {
	resolvers["Project"] = projectsResolver
}

const (
	ProjectCreate     = Action("ProjectCreate")
	ProjectsAll       = Action("ProjectsAll")
	ProjectFindById   = Action("ProjectFindById")
	ProjectAllUsers   = Action("ProjectAllUsers")
	ProjectAllTasks   = Action("ProjectAllTasks")
	ProjectAddUser    = Action("ProjectAddUser")
	ProjectDeleteUser = Action("ProjectDeleteUser")
)

func projectsResolver(action Action) (service ServiceFunc, err error) {
	switch action {
	case ProjectCreate:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			project, err := models.SafeCastToProject(data)
			if err != nil {
				return models.Project{}, err
			}
			return services.CreateProject(source, project)
		}
		return

	case ProjectsAll:
		service = func(source db.DataSource, _ interface{}) (interface{}, error) {
			return services.AllProjects(source)
		}
		return

	case ProjectFindById:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return models.Project{}, err
			}
			return services.FindProjectById(source, id)
		}
		return

	case ProjectAllUsers:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return services.AllUsersInProject(source, id)
			}
			return models.UsersList{}, err
		}
		return

	case ProjectAllTasks:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return models.TasksList{}, err
			}
			return services.AllTasksInProject(source, id)
		}
		return

	case ProjectAddUser:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			ids, err := models.SafeCastToProjectUser(data)
			if err != nil {
				return models.UsersList{}, err
			}
			return services.AddUserToProject(source, ids.ProjectId, ids.UserId)
		}
		return

	case ProjectDeleteUser:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			ids, err := models.SafeCastToProjectUser(data)
			if err != nil {
				return models.UsersList{}, err
			}
			return services.DeleteUserFromProject(source, ids.ProjectId, ids.UserId)
		}
		return

	}
	return nil, fmt.Errorf("can not find resolver with action: %v, in projects resolvers", action)
}
