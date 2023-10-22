package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lautaroDeLuca/go-rest-gorm/db"
	"github.com/lautaroDeLuca/go-rest-gorm/models"
	"github.com/lautaroDeLuca/go-rest-gorm/routes"
)

func main() {
	db.DbConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
