package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"

	"github.com/suryanadeva/digitalent-microservice/menu-service/database"
	"github.com/suryanadeva/digitalent-microservice/menu-service/utils"
)

type MenuHandler struct {
	Db *gorm.DB
}

func (handler *MenuHandler) AddMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	var menu database.Menu
	err = json.Unmarshal(body, &menu)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	err = menu.Insert(handler.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
	}

	utils.WrapAPISuccess(w, r, "success", http.StatusOK)

}

func (handler *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	menu := database.Menu{}

	menus, err := menu.GetAll(handler.Db)
	if err != nil {
		utils.WrapAPIError(w, r, "failed get menu:"+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(w, r, menus, http.StatusOK, "success")
}
