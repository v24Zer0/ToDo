package models

import (
	"encoding/json"
	"io"
)

type UserToken struct {
	UserID string `json:"userID"`
	Token  string `json:"token"`
}

func (t *UserToken) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(t)
}

func (t *UserToken) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(t)
}
