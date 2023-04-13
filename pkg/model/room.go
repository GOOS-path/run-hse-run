package model

type Room struct {
	Id       int64  `json:"id" db:"id"`
	Code     string `json:"code" db:"code"`
	CampusId int64  `json:"campus_id" db:"campus_id"`
}
