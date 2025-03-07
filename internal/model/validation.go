package model

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func (m *RegisterRequest) Validate() error {
	v := validator.New(validator.WithRequiredStructEnabled())

	err := v.Struct(m)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, fieldErr := range validationErrors {
				switch fieldErr.Tag() {
				case "email":
					return fmt.Errorf("Поле %s должно быть валидным email адресом\n", fieldErr.Field())
				default:
					return fmt.Errorf("Ошибка валидации для поля %s: %s\n", fieldErr.Field(), fieldErr.Tag())
				}
			}
		}
	}

	return nil
}

func (m *LoginRequest) Validate() error {
	v := validator.New(validator.WithRequiredStructEnabled())

	return v.Struct(m)
}
