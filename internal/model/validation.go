package model

import "github.com/go-playground/validator/v10"

func (m *RegisterRequest) Validate() error {
	v := validator.New(validator.WithRequiredStructEnabled())

	return v.Struct(m)
}

func (m *LoginRequest) Validate() error {
	v := validator.New(validator.WithRequiredStructEnabled())

	return v.Struct(m)
}
