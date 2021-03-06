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
	"github.com/v24Zer0/ToDo/user-service/database"
	"github.com/v24Zer0/ToDo/user-service/handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading environment variables")
	}

	port := os.Getenv("Port")

	router := mux.NewRouter()

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalln("Error connecting to database")
	}

	handler := handlers.NewHandler(db)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/user", handler.CreateUser)
	postRouter.HandleFunc("/login", handler.Login)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/user", handler.DeleteUser)

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
