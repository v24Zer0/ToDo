package utility

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
