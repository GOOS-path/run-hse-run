package repository

import (
	"Run_Hse_Run/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int64, error)
	GetUser(email string) (model.User, error)
}

type Friends interface {
	AddFriend(userIdFrom, userIdTo int64) error
	DeleteFriend(userIdFrom, userIdTo int64) error
	GetFriends(userId int64) ([]model.User, error)
}

type Users interface {
	GetUserById(userId int64) (model.User, error)
	GetUsersByNicknamePattern(nickname string) ([]model.User, error)
	RenameUser(userId int64, nickname string) error
	ChangeProfileImage(userId int64, image string) error
}

type Game interface {
	GetRoomByCodePattern(code string, campusId int64) ([]model.Room, error)
	GetEdge(startRoomId, endRoomId int64) (model.Edge, error)
	GetListOfEdges(startRoomId int64) ([]model.Edge, error)
	GetRoomById(roomId int64) (model.Room, error)
	AddCall(userIdFirst, userIdSecond, roomIdFirst int64) (model.Game, error)
	DeleteCall(userIdFirst, userIdSecond int64) error
	GetGame(gameId int64) (model.GameUsers, error)
	GetTime(gameId, userId int64) (model.Time, error)
	AddGame(userIdFirst, userIdSecond int64) (int64, error)
	AddTime(gameId, userId, time int64) error
	UpdateTime(gameId, userId, time int64) error
}

type Repository struct {
	Authorization
	Friends
	Users
	Game
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Friends:       NewFriendPostgres(db),
		Users:         NewUsersPostgres(db),
		Game:          NewGamePostgres(db),
	}
}
