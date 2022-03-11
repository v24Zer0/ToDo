package handlers

import (
	"log"
	"net/http"
)

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete item request")
}

func DeleteList(w http.ResponseWriter, r *http.Request) {

}
