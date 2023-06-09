package repository

import (
	"Run_Hse_Run/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user model.User) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (nickname, email, image, score) values ($1, $2, $3, $4) RETURNING id", usersTable)

	row := a.db.QueryRow(query, user.Nickname, user.Email, user.Image, user.Score)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthPostgres) GetUser(email string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1", usersTable)
	err := a.db.Get(&user, query, email)

	return user, err
}
