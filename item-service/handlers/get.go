package handlers

import (
	"log"
	"net/http"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	log.Println("Get item request")
}
