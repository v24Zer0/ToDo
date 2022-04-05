package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/v24Zer0/ToDO/user-service/authentication"
	"github.com/v24Zer0/ToDO/user-service/database"
	"github.com/v24Zer0/ToDO/user-service/models"
	"github.com/v24Zer0/ToDO/user-service/utility"
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

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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

	err = utility.ComparePassword(u.Password, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := authentication.GetToken(u.ID)
	response := models.UserResponse{
		ID:    u.ID,
		Token: token,
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
