package service

import (
	"Run_Hse_Run/pkg/mailer"
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/queue"
	"Run_Hse_Run/pkg/repository"
	"Run_Hse_Run/pkg/websocket"
	"net/http"
	"sync"
)

var (
	Mu    sync.Mutex
	Codes = make(map[string]int64)
)

type Sender interface {
	SendEmail(email string) error
}

type Authorization interface {
	CreateUser(user model.User) (int64, error)
	GetUser(email string) (model.User, error)
	GenerateToken(email string) (string, error)
	ParseToken(accessToken string) (int64, error)
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
	AddCall(userIdFirst, userIdSecond, roomIdFirst int64) (model.Game, error)
	DeleteCall(userIdFirst, userIdSecond int64) error
	AddUser(userId, roomId int64)
	Cancel(userId int64)
	SendGame(game model.Game) error
	UpgradeConnection(w http.ResponseWriter, r *http.Request)
	SendResult(gameId, userId, time int64)
	UpdateTime(gameId, userId, time int64) error
}

type Service struct {
	Sender
	Authorization
	Friends
	Users
	Game
}

func NewService(repo *repository.Repository, sender *mailer.Mailer,
	queue *queue.Queue, websocket *websocket.Server) *Service {
	return &Service{
		Sender:        NewSenderService(sender),
		Authorization: NewAuthService(repo),
		Friends:       NewFriendsService(repo),
		Users:         NewUsersService(repo),
		Game:          NewGameService(repo, queue, websocket),
	}
}
