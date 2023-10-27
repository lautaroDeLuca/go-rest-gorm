package routes

import (
	"github.com/gorilla/mux"
	"github.com/lautaroDeLuca/go-rest-gorm/handlers"
)

func DefineUserRoutes(r *mux.Router) {
	r.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", handlers.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
}
