package validation

import (
	"chetam/internal/model"
	"github.com/go-playground/validator/v10"
)

func ValidateRegisterRequest(req model.RegisterRequest) error {
	v := validator.New(validator.WithRequiredStructEnabled())

	return v.Struct(req)
}

func ValidateLoginRequest(req model.LoginRequest) error {
	v := validator.New(validator.WithRequiredStructEnabled())

	return v.Struct(req)
}
