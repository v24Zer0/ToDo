package models

import (
	"encoding/json"
	"io"
)

type List struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type Lists []*List

func (l *Lists) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(l)
}

func (l *Lists) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(l)
}
