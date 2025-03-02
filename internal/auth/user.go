package auth

import (
	"chetam/internal/model"
	"fmt"
	"log/slog"
)

func (a *Auth) CreateUser(req model.RegisterRequest) (string, error) {
	_, err := a.repo.CreateUser(req.Email, req.Login, req.Password)
	if err != nil {
		return "", err
	}

	token, err := a.generateJWT(req.Login)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Auth) UserToken(req model.LoginRequest) (string, error) {
	user, err := a.repo.FindUserByLogin(req.Login)
	if err != nil {
		a.lg.Warn("user not found",
			slog.String("error", err.Error()))
		return "", err
	} else if user.Password != req.Password {
		return "", fmt.Errorf("password incorrect")
	}

	token, err := a.generateJWT(user.Login)
	if err != nil {
		return "", err
	}
	return token, nil
}
