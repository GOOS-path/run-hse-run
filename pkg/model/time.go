package model

type Time struct {
	Id     int64 `json:"id" db:"id"`
	GameId int64 `json:"game_id" db:"game_id"`
	UserId int64 `json:"user_id" db:"user_id"`
	Time   int64 `json:"time" db:"time"`
}
