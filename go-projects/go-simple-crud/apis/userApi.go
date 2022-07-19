package apis

import (
	"main/go-projects/go-simple-crud/helpers"
	"main/go-projects/go-simple-crud/models"
	"net/http"
)

func FindUser(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "User Not Found!Pls check again.")
		return
	}
	user, err := models.FindUser(ids[0])
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
	}
	helpers.ResponseWithJSON(w, http.StatusOK, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := models.GetAllUser()
	if len(list) < 1 {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "Could not find any user!")
	}
	helpers.ResponseWithJSON(w, http.StatusOK, list)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// var user entities.User
	// err := json.NewDecoder(r.Body).Decode(&user)
}

// func UpdateUser()
