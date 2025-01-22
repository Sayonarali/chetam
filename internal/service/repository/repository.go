package repository

import (
	dbClient "chetam/internal/db/client"
	"chetam/internal/models"
	"github.com/Masterminds/squirrel"
)

type Keeper struct {
	chetamFetcher *dbClient.ChetamFetcher
}

var sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func NewRepositoryKeeper(chf *dbClient.ChetamFetcher) Keeper {
	return Keeper{
		chetamFetcher: chf,
	}
}

func (k Keeper) FindUserByLogin(login string) (models.User, error) {
	user := models.User{}

	selectQuery, args, err := sq.Select("id, login, email, password").From("users").Where(`"login" = ?`, login).ToSql()
	if err != nil {
		return user, err
	}

	err = k.chetamFetcher.C.QueryRow(selectQuery, args...).Scan(&user.Id, &user.Login, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (k Keeper) CreateUser(email, login, password string) (models.User, error) {
	user := models.User{}

	createQuery, args, err := sq.Insert("users").Columns("login", "email", "password").Values(login, email, password).ToSql()
	if err != nil {
		return user, err
	}

	_, err = k.chetamFetcher.C.Exec(createQuery, args...)
	if err != nil {
		return user, err
	}

	user = models.User{
		Login:    login,
		Email:    email,
		Password: password,
	}

	return user, nil
}
