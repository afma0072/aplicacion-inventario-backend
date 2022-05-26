package routes

import (
	"github.com/afma0072/aplicacion-inventario-backend/controllers"
	"github.com/gorilla/mux"
)

func SetCategoriaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/categoria/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")

}
