package model

type Friendship struct {
	Id         int64 `json:"id" db:"id"`
	UserIdFrom int64 `json:"user_id_from" db:"user_id1"`
	UserIdTo   int64 `json:"user_id_to" db:"user_id2"`
}
