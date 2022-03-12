package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v24Zer0/ToDo/item-service/handlers"
)

func main() {
	router := mux.NewRouter()

	itemHandler := handlers.NewHandler()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/item", itemHandler.GetItems)

	log.Println("Server starting on port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
