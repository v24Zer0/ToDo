package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/v24Zer0/ToDo/item-service/authentication"
	"github.com/v24Zer0/ToDo/item-service/database"
	"github.com/v24Zer0/ToDo/item-service/handlers"
)

func main() {
	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading environment variables")
	}

	router := mux.NewRouter()

	// authentication middleware for each route
	router.Use(authentication.AuthMiddleware)

	// create new database connection
	db, err := database.NewDatabase()
	if err != nil {
		log.Println("Error connecting to database")
	}

	itemHandler := handlers.NewHandler(db)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/items/{id:[a-zA-Z0-9]{27}}", itemHandler.GetItems)
	getRouter.HandleFunc("/lists/{id:[a-zA-Z0-9]{27}}", itemHandler.GetLists)
	getRouter.HandleFunc("/id", itemHandler.GetID)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/item", itemHandler.CreateItem)
	postRouter.HandleFunc("/list", itemHandler.CreateList)

	updateRouter := router.Methods(http.MethodPut).Subrouter()
	updateRouter.HandleFunc("/item", itemHandler.UpdateItem)
	updateRouter.HandleFunc("/list", itemHandler.UpdateList)

	// deleteRouter := router.Methods(http.MethodDelete).Subrouter()

	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}

	// start http server
	go func() {
		log.Println("Server starting on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			log.Println("Server shutdown")
		}
	}()

	// trap sigterm or interupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	os.Clearenv()
}
