package handlers

import (
	"chetam/internal/model"
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
		var e model.Error

		if err := c.Bind(&req); err != nil {
			e = model.Error{
				Error: model.ErrInvalidJson,
			}

			logger.Error(e.Error, slog.String("error", err.Error()))
			return c.JSON(http.StatusBadRequest, e)
		}

		err := req.Validate()
		if err != nil {
			e = model.Error{
				Error: err.Error(),
			}

			logger.Error(e.Error, slog.String("error", err.Error()))
			return c.JSON(http.StatusBadRequest, e)
		}

		token, err := auth.CreateUser(req)
		if err != nil {
			e = model.Error{
				Error: model.ErrNotUniqueUser,
			}
			logger.Error(e.Error, slog.String("error", err.Error()))
			return c.JSON(http.StatusBadRequest, e)
		}
		var resp model.RegisterResponse

		resp.Token = token

		return c.JSON(http.StatusOK, resp)
	}
}

func Login(logger *slog.Logger, auth AuthInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.LoginRequest
		var e model.Error

		if err := c.Bind(&req); err != nil {
			e = model.Error{
				Error: model.ErrInvalidJson,
			}

			logger.Error(e.Error, slog.String("error", err.Error()))
			return c.JSON(http.StatusBadRequest, e)
		}

		err := req.Validate()
		if err != nil {
			e = model.Error{
				Error: err.Error(),
			}

			logger.Error(e.Error, slog.String("error", err.Error()))
			return c.JSON(http.StatusBadRequest, e)
		}

		token, err := auth.UserToken(req)
		if err != nil {
			e = model.Error{
				Error: model.ErrInternal,
			}
			logger.Error(e.Error, slog.String("error", err.Error()))
			return c.JSON(http.StatusBadRequest, e)
		}
		var resp model.LoginResponse

		resp.Token = token

		return c.JSON(http.StatusOK, resp)
	}
}
