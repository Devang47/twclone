package models

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Photo    string `json:"photo"`
	Uid      string `json:"uid"`
}
