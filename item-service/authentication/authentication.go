package authentication

import (
	"net/http"
	"os"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authURL := os.Getenv("AuthService")
		next.ServeHTTP(w, r)
	})
}
