package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/afma0072/aplicacion-inventario-backend/commons"
	"github.com/afma0072/aplicacion-inventario-backend/models"
	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	categorias := []models.Categoria{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&categorias)
	json, _ := json.Marshal(categorias)
	commons.SendReponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	categoria := models.Categoria{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&categoria, id)

	if categoria.ID > 0 {
		json, _ := json.Marshal(categoria)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func Save(writer http.ResponseWriter, request *http.Request) {
	categoria := models.Categoria{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&categoria)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&categoria).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(categoria)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	categoria := models.Categoria{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&categoria, id)

	if categoria.ID > 0 {
		db.Delete(categoria)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
