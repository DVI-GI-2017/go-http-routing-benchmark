package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

func init() {
	defaultRoutes = append(defaultRoutes, tasksRoutes...)
}

var tasksRoutes = []Route{
	{
		Name:    "Add task to project.",
		Pattern: "/projects/{hex:project_id}/tasks",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.AddTaskToProject),
	},
	{
		Name:    "All tasks in project",
		Pattern: "/projects/{hex:project_id}/tasks",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllTasksInProject),
	},
	{
		Name:    "Get task by id",
		Pattern: "/tasks/{hex:id}",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetTaskById),
	},
}
