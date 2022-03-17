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

func GetLists(id string) Lists {
	lists := Lists{}
	for _, list := range userLists {
		if list.UserID == id {
			lists = append(lists, list)
		}
	}
	return lists
}

var userLists Lists = Lists{
	{
		ID:     "0",
		Name:   "Life",
		UserID: "0",
	},
	{
		ID:     "1",
		Name:   "Ideas",
		UserID: "0",
	},
	{
		ID:     "2",
		Name:   "Projects",
		UserID: "1",
	},
	{
		ID:     "3",
		Name:   "Other",
		UserID: "1",
	},
	{
		ID:     "4",
		Name:   "Work",
		UserID: "2",
	},
	{
		ID:     "5",
		Name:   "Home",
		UserID: "3",
	},
}
