package handlers

import (
	"chetam/internal/model"
	"chetam/internal/validation"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type AuthInterface interface {
	CreateUser(login string) (*model.User, error)
	FindUserByLogin(login string) (*model.User, error)
	Login(login, password string) (string, error)
}

func Register(logger *slog.Logger, auth AuthInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.RegisterRequest

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		err := validation.ValidateAuthRequest(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		user, err := auth.FindUserByLogin(req.Login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		user, err := auth.CreateUser(req.Email, req.Login, req.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return nil
	}
}

func Login(logger *slog.Logger, auth AuthInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusBadGateway, "fsddsfs")
	}
}
