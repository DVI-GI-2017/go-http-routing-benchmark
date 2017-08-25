package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

func init() {
	defaultRoutes = append(defaultRoutes, projectRoutes...)
}

var projectRoutes = []Route{
	{
		Name:    "Creates project",
		Pattern: "/projects",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.CreateProject),
	},
	{
		Name:    "Get all projects",
		Pattern: "/projects",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllProjects),
	},
	{
		Name:    "Get project by id",
		Pattern: "/projects/:id",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetProjectById),
	},
	{
		Name:    "Get all users from project with given id",
		Pattern: "/projects/:id/users",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetAllUsersFromProject),
	},
	{
		Name:    "Get all tasks from project with given id",
		Pattern: "/projects/:id/tasks",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetAllTasksFromProject),
	},
	{
		Name:    "Add user to project with given id",
		Pattern: "/projects/:id/users",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.AddUserToProject),
	},
	{
		Name:    "Delete user from project with given id",
		Pattern: "/projects/:id/users/delete",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.DeleteUserFromProject),
	},
}
