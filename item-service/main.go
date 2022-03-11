package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v24Zer0/ToDo/item-service/handlers"
)

func main() {
	router := mux.NewRouter()

	// handler := handlers.NewHandler()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/item", handlers.GetItems)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/item", handlers.CreateItem)

	updateRouter := router.Methods(http.MethodPut).Subrouter()
	updateRouter.HandleFunc("/item", handlers.UpdateItem)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/item", handlers.DeleteItem)

	log.Println("Server starting on port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
