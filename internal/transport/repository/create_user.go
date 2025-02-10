package repository

import (
	"chetam/internal/model"
	"context"
	"time"
)

func (r *Repository) CreateUser(email, login, password string) (model.User, error) {
	querry := qb.Insert("users").
		Columns(
			"login",
			"email",
			"password",
		).
		Values(
			login,
			email,
			password,
		)

	sql, args, err := querry.ToSql()
	if err != nil {
		return model.User{}, err
	}

	_, err = r.db.Exec(context.Background(), sql, args...)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Login:     login,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}
