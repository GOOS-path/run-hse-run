package service

import (
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/repository"
	"errors"
	"fmt"
	"regexp"
)

type UsersService struct {
	repo *repository.Repository
}

func (u *UsersService) ChangeProfileImage(userId int64, image string) error {
	return u.repo.ChangeProfileImage(userId, image)
}

func (u *UsersService) RenameUser(userId int64, nickname string) error {
	if nickname == "" {
		return errors.New(NicknameError)
	}

	if 15 < len(nickname) {
		return errors.New(NicknameError)
	}

	expr := fmt.Sprintf("^[a-zA-Z0-9_]{%d}", len(nickname))
	validUser, err := regexp.Compile(expr)
	if err != nil {
		return errors.New(NicknameError)
	}

	if !validUser.MatchString(nickname) {
		return errors.New(NicknameError)
	}

	return u.repo.RenameUser(userId, nickname)
}

func (u *UsersService) GetUsersByNicknamePattern(nickname string) ([]model.User, error) {
	return u.repo.GetUsersByNicknamePattern(nickname)
}

func (u *UsersService) GetUserById(userId int64) (model.User, error) {
	return u.repo.GetUserById(userId)
}

func (u *UsersService) GetUsers() ([]model.User, error) {
	return u.repo.GetUsers()
}

func NewUsersService(repo *repository.Repository) *UsersService {
	return &UsersService{repo: repo}
}
