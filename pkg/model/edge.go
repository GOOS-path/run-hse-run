package model

type Edge struct {
	Id          int64   `json:"id" db:"id"`
	StartRoomId int64   `json:"start_room_id" db:"start_room_id"`
	EndRoomId   int64   `json:"end_room_id" db:"end_room_id"`
	Cost        float64 `json:"cost" db:"cost"`
	CampusId    int64   `json:"campus_id" db:"campus_id"`
}
