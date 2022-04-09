package main

import (
	"context"
	"fmt"
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

	port := os.Getenv("Port")

	router := mux.NewRouter()

	// authentication middleware for each route
	router.Use(authentication.AuthMiddleware)

	// create new database connection
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalln("Error connecting to database")
	}

	handler := handlers.NewHandler(db)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/items/{id:[a-zA-Z0-9]{27}}", handler.GetItems)
	getRouter.HandleFunc("/lists/{id:[a-zA-Z0-9]{27}}", handler.GetLists)
	getRouter.HandleFunc("/id", handler.GetID)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/item", handler.CreateItem)
	postRouter.HandleFunc("/list", handler.CreateList)

	updateRouter := router.Methods(http.MethodPut).Subrouter()
	updateRouter.HandleFunc("/item", handler.UpdateItem)
	updateRouter.HandleFunc("/list", handler.UpdateList)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/item/{id:[a-zA-Z0-9]{27}}", handler.DeleteItem)
	deleteRouter.HandleFunc("/list/{id:[a-zA-Z0-9]{27}}", handler.DeleteList)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// start http server
	go func() {
		log.Printf("Server starting on port %s", port)

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
