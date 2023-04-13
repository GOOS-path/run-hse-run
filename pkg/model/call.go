package model

type Call struct {
	Id           int64 `json:"id" db:"id"`
	UserIdFirst  int64 `json:"user_id_first" db:"user_id_first"`
	RoomIdFirst  int64 `json:"room_id_first" db:"room_id_first"`
	UserIdSecond int64 `json:"user_id_second" db:"user_id_second"`
}
