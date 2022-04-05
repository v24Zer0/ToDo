package models

import (
	"encoding/json"
	"io"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserNoPassword struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (u *User) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}

func (u *UserNoPassword) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(u)
}

func (u *UserNoPassword) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}

func (u *UserResponse) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(u)
}
