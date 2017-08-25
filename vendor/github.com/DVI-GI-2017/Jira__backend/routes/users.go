package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

func init() {
	defaultRoutes = append(defaultRoutes, usersRoutes...)
}

var usersRoutes = []Route{
	{
		Name:    "Get user by id",
		Pattern: "/users/hex:id",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetUserById),
	},
	{
		Name:    "All users route",
		Pattern: "/users",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllUsers),
	},
	{
		Name:    "All user's projects",
		Pattern: "/users/hex:id/projects",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetAllProjectsFromUser),
	},
	{
		Name:    "Check current user",
		Pattern: "/cur-user",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.JsonNullHandler),
	},
}
