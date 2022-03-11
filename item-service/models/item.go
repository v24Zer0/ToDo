package models

type Item struct {
	Id       string `json:"id"`
	Task     string `json:"task"`
	Priority int    `json:"priority"`
	ListId   string `json:"listId"`
}
