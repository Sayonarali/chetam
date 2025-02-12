package repository

import (
	"chetam/internal/model"
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) FindUserByLogin(login string) (model.User, error) {
	querry := qb.Select(
		"id",
		"login",
		"email",
		"password",
	).From("users").
		Where(sq.Eq{
			"login": login,
		})

	sql, args, err := querry.ToSql()
	if err != nil {
		return model.User{}, err
	}

	user := model.User{}

	err = r.db.QueryRow(context.Background(), sql, args...).Scan(&user.Id, &user.Login, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
