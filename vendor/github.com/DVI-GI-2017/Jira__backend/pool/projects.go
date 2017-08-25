package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/projects"
)

func init() {
	resolvers["Project"] = projectsResolver
}

const (
	ProjectCreate     = Action("ProjectCreate")
	ProjectExists     = Action("ProjectExists")
	ProjectsAll       = Action("ProjectsAll")
	ProjectFindById   = Action("ProjectFindById")
	ProjectAllUsers   = Action("ProjectAllUsers")
	ProjectAllTasks   = Action("ProjectAllTasks")
	ProjectAddUser    = Action("ProjectAddUser")
	ProjectDeleteUser = Action("ProjectDeleteUser")
	ProjectUserExists = Action("ProjectUserExists")
)

func projectsResolver(action Action) (service ServiceFunc) {
	switch action {
	case ProjectCreate:
		service = func(source db.DataSource, project interface{}) (interface{}, error) {
			return projects.CreateProject(source, project.(models.Project))
		}
		return

	case ProjectExists:
		service = func(source db.DataSource, project interface{}) (interface{}, error) {
			return projects.CheckProjectExists(source, project.(models.Project))
		}
		return

	case ProjectsAll:
		service = func(source db.DataSource, _ interface{}) (interface{}, error) {
			return projects.AllProjects(source)
		}
		return

	case ProjectFindById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return projects.FindProjectById(source, id.(models.RequiredId))
		}
		return

	case ProjectAllUsers:
		service = func(source db.DataSource, id interface{}) (result interface{}, err error) {
			return projects.AllUsersInProject(source, id.(models.RequiredId))
		}
		return

	case ProjectAllTasks:
		service = func(source db.DataSource, id interface{}) (result interface{}, err error) {
			return projects.AllTasksInProject(source, id.(models.RequiredId))
		}
		return

	case ProjectAddUser:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			ids := data.(models.ProjectUser)
			return projects.AddUserToProject(source, ids.ProjectId, ids.UserId)
		}
		return

	case ProjectDeleteUser:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			ids := data.(models.ProjectUser)
			return projects.DeleteUserFromProject(source, ids.ProjectId, ids.UserId)
		}
		return

	case ProjectUserExists:
		service = func(source db.DataSource, data interface{}) (result interface{}, err error) {
			ids := data.(models.ProjectUser)
			return projects.CheckUserInProject(source, ids.UserId, ids.ProjectId)
		}
		return

	default:
		log.Panicf("can not find resolver with action: %v, in projects resolvers", action)
		return
	}
}
