package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lautaroDeLuca/go-rest-gorm/db"
	"github.com/lautaroDeLuca/go-rest-gorm/models"
	"github.com/lautaroDeLuca/go-rest-gorm/routes"
)

func main() {
	// Connect to db, migrate tables
	db.DbConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	// Get the Router to set the route handlers
	r := mux.NewRouter()

	// Initiate routes
	routes.DefineRoutes(r)

	http.ListenAndServe(":3000", r)
}
