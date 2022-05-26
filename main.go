package main

import (
	"log"
	"net/http"

	"github.com/afma0072/aplicacion-inventario-backend/commons"
	"github.com/afma0072/aplicacion-inventario-backend/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetCategoriaRoutes(router)
	routes.SetProductoRoutes(router)
	commons.EnableCORS(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")
	log.Println(server.ListenAndServe())
}
