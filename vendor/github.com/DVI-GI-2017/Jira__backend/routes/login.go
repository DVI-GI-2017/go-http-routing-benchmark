package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

func init() {
	defaultRoutes = append(defaultRoutes, loginRoutes...)
}

var loginRoutes = []Route{
	{
		Name:    "Sign-up route",
		Pattern: "/signup",
		Method:  http.MethodPost,
		Handler: handlers.RegisterUser,
	},
	{
		Name:    "Sign-in route",
		Pattern: "/signin",
		Method:  http.MethodPost,
		Handler: handlers.Login,
	},
}
