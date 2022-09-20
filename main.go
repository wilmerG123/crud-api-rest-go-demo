package main

import (
	"crud-api-rest-go-demo/commons"
	"crud-api-rest-go-demo/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)
	commons.EnableCORS(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 8080")
	log.Println(server.ListenAndServe())
}
