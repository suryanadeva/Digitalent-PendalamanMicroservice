package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"

	"github.com/suryanadeva/digitalent-microservice/auth-service/database"
	"github.com/suryanadeva/digitalent-microservice/utils"
)

type AuthDB struct {
	Db *gorm.DB
}

func ValidateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	if authToken != "respecker" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	utils.WrapAPISuccess(w, r, "success", 200)
}

//TODO Buat signup

func (db *AuthDB) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "cannot read body", http.StatusBadRequest)
		return
	}

	var signup database.Auth

	err = json.Unmarshal(body, &signup)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}

	signup.Token = utils.IdGenerator()
	err = signup.SignUp(db.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusBadRequest)

	}

	utils.WrapAPIError(w, r, "Success", http.StatusOK)
	return

	//TODO Buat login
}

func (db *AuthDB) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

}
