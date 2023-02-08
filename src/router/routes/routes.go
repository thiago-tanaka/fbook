package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"webapp/src/middlewares"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, usersRoutes...)
	routes = append(routes, usersRoutes...)
	routes = append(routes, mainPageRoute)

	for _, route := range routes {
		if route.RequiresAuthentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
