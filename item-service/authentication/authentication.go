package authentication

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/v24Zer0/ToDo/item-service/models"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authURL := os.Getenv("AuthService")

		token := r.Header.Get("Token")
		usertoken := models.UserToken{Token: token}

		var payload bytes.Buffer
		err := usertoken.Encode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		resp, err := http.Post(authURL, "application/json", &payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		defer resp.Body.Close()

		err = usertoken.Decode(resp.Body)
		if err != nil {
			return
		}

		log.Println(usertoken.Token)
		r.Header.Add("User_ID", usertoken.UserID)

		next.ServeHTTP(w, r)
	})
}
