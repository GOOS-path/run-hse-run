package model

type User struct {
	Id       int64  `json:"id" db:"id"`
	Nickname string `json:"nickname" db:"nickname"`
	Email    string `json:"email" db:"email"`
	Image    string `json:"image" db:"image"`
	Score    int64  `json:"score" db:"score"`
}
