package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lautaroDeLuca/go-rest-gorm/db"
	"github.com/lautaroDeLuca/go-rest-gorm/models"
	"github.com/lautaroDeLuca/go-rest-gorm/validator"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	validatorErr := validator.Validate.Struct(user)
	if validatorErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validatorErr.Error()))
		return
	}

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	deleteResult := db.DB.Delete(&user, params["id"])

	if deleteResult.Error != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(deleteResult.Error.Error()))
		return
	}
	if deleteResult.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}
