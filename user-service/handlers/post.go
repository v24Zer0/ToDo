package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/v24Zer0/ToDo/user-service/authentication"
	"github.com/v24Zer0/ToDo/user-service/database"
	"github.com/v24Zer0/ToDo/user-service/models"
	"github.com/v24Zer0/ToDo/user-service/utility"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Post request - User")

	var user models.User
	err := user.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateUserRequest(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.CreateUser(h.db, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := make(chan bool)

	go authentication.CreateToken(c, user.ID)

	success := <-c
	if !success {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Post request - Login")

	var user models.User
	err := user.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateUserRequest(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := database.RetrieveUser(h.db, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c1 := make(chan error)
	c2 := make(chan authentication.AuthToken)

	go utility.ComparePassword(c1, u.Password, user.Password)
	go authentication.GetToken(c2, u.ID)

	err = <-c1
	token := <-c2

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if token.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.UserResponse{
		ID:    u.ID,
		Token: token.Token,
	}

	w.Header().Add("Content-Type", "application/json")
	response.Encode(w)
}

func validateUserRequest(u *models.User) error {
	if u.Username == "" {
		return errors.New("missing username")
	}

	if u.Password == "" {
		return errors.New("missing password")
	}
	return nil
}
