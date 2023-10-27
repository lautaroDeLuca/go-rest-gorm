package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lautaroDeLuca/go-rest-gorm/db"
	"github.com/lautaroDeLuca/go-rest-gorm/models"
	"github.com/lautaroDeLuca/go-rest-gorm/validator"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	params := mux.Vars(r)
	db.DB.First(&tasks, params["id"])

	if tasks.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&tasks)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	if err := validator.Validate.Struct(task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if createdTask := db.DB.Create(&task); createdTask.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(createdTask.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	if result := db.DB.First(&task, params["id"]); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(result.Error.Error()))
		return
	}

	deleteTask := db.DB.Delete(&task, params["id"])

	if deleteTask.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(deleteTask.Error.Error()))
		return
	}
	if deleteTask.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}
