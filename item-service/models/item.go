package models

import (
	"encoding/json"
	"io"
)

type Item struct {
	ID       string `json:"id"`
	Task     string `json:"task"`
	Priority int    `json:"priority"`
	ListID   string `json:"list_id"`
}

func (item *Item) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(item)
}

type Items []*Item

func (items *Items) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(items)
}
