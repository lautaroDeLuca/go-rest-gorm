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

type Data struct {
	clientId  string
	Seller    int32
	userId    int8
	monday    bool
	tuesday   bool
	wednesday bool
	thursday  bool
	friday    bool
	saturday  bool
	sunday    bool
}
