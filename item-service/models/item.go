package models

import (
	"encoding/json"
	"io"
)

type Item struct {
	Id       string `json:"id"`
	Task     string `json:"task"`
	Priority int    `json:"priority"`
	ListId   string `json:"listId"`
}

func (item *Item) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(item)
}

type Items []*Item

func (items Items) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(items)
}

func GetItems() Items {
	return itemList
}

var itemList = Items{
	{
		Id:       "0",
		Task:     "Update list",
		Priority: 0,
		ListId:   "0",
	},
	{
		Id:       "1",
		Task:     "Get groceries",
		Priority: 0,
		ListId:   "0",
	},
	{
		Id:       "2",
		Task:     "Wash car",
		Priority: 0,
		ListId:   "0",
	},
	{
		Id:       "3",
		Task:     "Complete project",
		Priority: 0,
		ListId:   "0",
	},
}
