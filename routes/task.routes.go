package routes

import (
	"github.com/gorilla/mux"
	"github.com/lautaroDeLuca/go-rest-gorm/handlers"
)

func DefineTaskRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", handlers.PostTaskHandler).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
}
