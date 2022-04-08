package models

type UserToken struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
