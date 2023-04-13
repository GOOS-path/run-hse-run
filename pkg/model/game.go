package model

type Game struct {
	UserIdFirst, RoomIdFirst   int64
	UserIdSecond, RoomIdSecond int64
}

type GameUsers struct {
	Id           int64 `json:"id" db:"id"`
	UserIdFirst  int64 `json:"user_id_first" db:"user_id_first"`
	UserIdSecond int64 `json:"user_id_second" db:"user_id_second"`
}
