package authentication

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/v24Zer0/ToDo/user-service/models"
)

type AuthToken struct {
	Token string
	Error error
}

type UserValidation struct {
	UserID string
	Error  error
}

func GetToken(c chan AuthToken, id string) {
	authURL := os.Getenv("AuthService")
	url := fmt.Sprintf("%s/%s", authURL, id)

	resp, err := http.Get(url)
	if err != nil {
		c <- AuthToken{Error: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c <- AuthToken{Error: errors.New("error fetching token")}
		return
	}

	var userToken models.UserToken
	userToken.Decode(resp.Body)

	c <- AuthToken{Token: userToken.Token}
}

func CreateToken(c chan bool, id string) {
	authURL := os.Getenv("AuthService")

	var payload bytes.Buffer
	userToken := models.UserToken{UserID: id}

	err := userToken.Encode(&payload)
	if err != nil {
		c <- false
		return
	}

	resp, err := http.Post(authURL, "application/json", &payload)
	if err != nil {
		c <- false
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		c <- false
		return
	}

	c <- true
}

func ValidateToken(c chan UserValidation, token string) {
	authService := os.Getenv("AuthService")
	authURL := fmt.Sprintf("%s/%s", authService, "/validate")

	userToken := models.UserToken{Token: token}
	var payload bytes.Buffer

	err := userToken.Encode(&payload)
	if err != nil {
		c <- UserValidation{Error: err}
		return
	}

	resp, err := http.Post(authURL, "application/json", &payload)
	if err != nil {
		c <- UserValidation{Error: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c <- UserValidation{Error: errors.New("not authenticated")}
		return
	}

	err = userToken.Decode(resp.Body)
	if err != nil {
		c <- UserValidation{Error: err}
		return
	}

	c <- UserValidation{UserID: userToken.UserID}
}
