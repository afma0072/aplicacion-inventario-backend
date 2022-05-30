package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/afma0072/aplicacion-inventario-backend/commons"
	"github.com/afma0072/aplicacion-inventario-backend/models"
	"github.com/gorilla/mux"
)

func GetAllTransaccion(writer http.ResponseWriter, request *http.Request) {
	transaccion := []models.Transaccion{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&transaccion)
	json, _ := json.Marshal(transaccion)
	commons.SendReponse(writer, http.StatusOK, json)
}

func GetTransaccion(writer http.ResponseWriter, request *http.Request) {
	transaccion := models.Transaccion{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&transaccion, id)

	if transaccion.ID > 0 {
		json, _ := json.Marshal(transaccion)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func SaveTransaccion(writer http.ResponseWriter, request *http.Request) {
	transaccion := models.Transaccion{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&transaccion)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&transaccion).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(transaccion)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func DeleteTransaccion(writer http.ResponseWriter, request *http.Request) {
	transaccion := models.Transaccion{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&transaccion, id)

	if transaccion.ID > 0 {
		db.Delete(transaccion)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
