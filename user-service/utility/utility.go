package utility

import "golang.org/x/crypto/bcrypt"

func ComparePassword(c chan error, hash string, password string) {
	c <- bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
