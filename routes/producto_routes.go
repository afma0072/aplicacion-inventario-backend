package routes

import (
	"github.com/afma0072/aplicacion-inventario-backend/controllers"
	"github.com/gorilla/mux"
)

func SetProductoRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/producto/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAllProducto).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveProducto).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteProducto).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetProducto).Methods("GET")

}
