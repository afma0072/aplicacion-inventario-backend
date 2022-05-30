package routes

import (
	"github.com/afma0072/aplicacion-inventario-backend/controllers"
	"github.com/gorilla/mux"
)

func SetTransaccionRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/transaccion/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAllTransaccion).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveTransaccion).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteTransaccion).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetTransaccion).Methods("GET")

}
