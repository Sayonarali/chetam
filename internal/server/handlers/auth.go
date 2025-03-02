package handlers

import (
	"chetam/internal/model"
	"chetam/internal/validation"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type AuthInterface interface {
	CreateUser(req model.RegisterRequest) (string, error)
	UserToken(req model.LoginRequest) (string, error)
}

func Register(logger *slog.Logger, auth AuthInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.RegisterRequest

		if err := c.Bind(&req); err != nil {

			logger.Error("error binding request", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		err := validation.ValidateRegisterRequest(req)
		if err != nil {
			logger.Error("error validating request", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		token, err := auth.CreateUser(req)
		if err != nil {
			logger.Error("error creating token", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		var resp model.RegisterResponse

		resp.Token = token

		return c.JSON(http.StatusOK, resp)
	}
}

func Login(logger *slog.Logger, auth AuthInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.LoginRequest

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		err := validation.ValidateLoginRequest(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		token, err := auth.UserToken(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		var resp model.LoginResponse

		resp.Token = token

		return c.JSON(http.StatusOK, resp)
	}
}
