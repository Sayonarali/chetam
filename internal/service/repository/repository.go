package repository

import (
	dbClient "chetam/internal/db/client"
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

func (k Keeper) FindUserByLogin(username string) (string, error) {
	var id int
	var login, email, password string

	selectBuilder := sq.Select("id, login, email, password").From("users").Where(`"login" = ?`, username)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return "", err
	}

	err = k.chetamFetcher.C.QueryRow(query, args...).Scan(&id, &login, &email, &password)
	if err != nil {
		return "", err
	}

	return password, nil
}
