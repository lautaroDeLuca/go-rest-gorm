package routes

import "github.com/gorilla/mux"

func DefineRoutes(r *mux.Router) {
	DefineUserRoutes(r)
	DefineTaskRoutes(r)
}
