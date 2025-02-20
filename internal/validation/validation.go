package validation

import (
	"chetam/internal/model"
	"github.com/go-playground/validator/v10"
)

func ValidateAuthRequest(req model.RegisterRequest) error {
	v := validator.New(validator.WithRequiredStructEnabled())

	return v.Struct(req)
}
