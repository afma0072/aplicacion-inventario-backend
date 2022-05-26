package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/afma0072/aplicacion-inventario-backend/commons"
	"github.com/afma0072/aplicacion-inventario-backend/models"
	"github.com/gorilla/mux"
)

func GetAllProducto(writer http.ResponseWriter, request *http.Request) {
	producto := []models.Producto{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&producto)
	json, _ := json.Marshal(producto)
	commons.SendReponse(writer, http.StatusOK, json)
}

func GetProducto(writer http.ResponseWriter, request *http.Request) {
	producto := models.Producto{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&producto, id)

	if producto.ID > 0 {
		json, _ := json.Marshal(producto)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func SaveProducto(writer http.ResponseWriter, request *http.Request) {
	producto := models.Producto{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&producto)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&producto).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(producto)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func DeleteProducto(writer http.ResponseWriter, request *http.Request) {
	producto := models.Producto{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&producto, id)

	if producto.ID > 0 {
		db.Delete(producto)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
