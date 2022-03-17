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

func GetItems(id string) Items {
	items := Items{}
	for _, item := range itemList {
		if item.ListID == id {
			items = append(items, item)
		}
	}
	return items
}

var itemList = Items{
	{
		ID:       "0",
		Task:     "Update list",
		Priority: 0,
		ListID:   "0",
	},
	{
		ID:       "1",
		Task:     "Get groceries",
		Priority: 0,
		ListID:   "0",
	},
	{
		ID:       "2",
		Task:     "Wash car",
		Priority: 0,
		ListID:   "0",
	},
	{
		ID:       "3",
		Task:     "Complete project",
		Priority: 0,
		ListID:   "0",
	},
	{
		ID:       "4",
		Task:     "Test API",
		Priority: 0,
		ListID:   "1",
	},
	{
		ID:       "5",
		Task:     "Get things",
		Priority: 0,
		ListID:   "1",
	},
}
