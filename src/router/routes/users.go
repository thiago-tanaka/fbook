package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:                    "/create-account",
		Method:                 http.MethodGet,
		Function:               controllers.LoadCreateAccountPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
}
