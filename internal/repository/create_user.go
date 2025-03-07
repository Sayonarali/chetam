package repository

import (
	"chetam/internal/model"
	"context"
)

func (r *Repository) CreateUser(email, login, password string) (model.User, error) {
	query := qb.Insert("users").
		Columns(
			"login",
			"email",
			"password",
		).
		Values(
			login,
			email,
			password,
		).
		Suffix("RETURNING id, login, email, password")

	sql, args, err := query.ToSql()
	if err != nil {
		return model.User{}, err
	}

	user := model.User{}
	err = r.db.QueryRow(context.Background(), sql, args...).Scan(&user.Id, &user.Login, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
