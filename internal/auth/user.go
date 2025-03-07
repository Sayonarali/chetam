package auth

import (
	"chetam/internal/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func (a *Auth) CreateUser(req model.RegisterRequest) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user, err := a.repo.CreateUser(req.Email, req.Login, string(passHash))
	if err != nil {
		return "", err
	}

	token, err := a.generateJWT(user.Id)
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
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", fmt.Errorf("password incorrect")
	}

	token, err := a.generateJWT(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}
