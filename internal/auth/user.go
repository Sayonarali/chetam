package auth

import "chetam/internal/model"

func (a *Auth) CreateUser(email, login, password string) (string, error) {
	user, err := a.repo.CreateUser(email, login, password)
	if err != nil {
		return "", err
	}

	token, err := a.generateJWT(user.Login)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Auth) FindUserByLogin(login string) (model.User, error) {
	user, err := a.repo.FindUserByLogin(login)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
