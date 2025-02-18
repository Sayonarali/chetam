package auth

func (a *Auth) User(email, login, password string) (string, error) {
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
