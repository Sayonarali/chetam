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
		var errorMsg string

		if err := c.Bind(&req); err != nil {
			errorMsg = err.Error()
			logger.Error("binding request", slog.String("error", errorMsg))
			return c.JSON(http.StatusBadRequest, model.Error{
				Msg: errorMsg,
			})
		}

		err := req.Validate()
		if err != nil {
			errorMsg = err.Error()
			logger.Error("validating request", slog.String("error", errorMsg))
			return c.JSON(http.StatusBadRequest, model.Error{
				Msg: errorMsg,
			})
		}

		token, err := auth.CreateUser(req)
		if err != nil {
			errorMsg = err.Error()
			logger.Error("creating user", slog.String("error", errorMsg))
			return c.JSON(http.StatusInternalServerError, model.Error{
				Msg: errorMsg,
			})
		}
		var resp model.RegisterResponse

		resp.Token = token

		return c.JSON(http.StatusOK, resp)
	}
}

func Login(logger *slog.Logger, auth AuthInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.LoginRequest
		var errorMsg string

		if err := c.Bind(&req); err != nil {
			errorMsg = err.Error()
			logger.Error("binding request", slog.String("error", errorMsg))
			return c.JSON(http.StatusBadRequest, model.Error{
				Msg: errorMsg,
			})
		}

		err := req.Validate()
		if err != nil {
			errorMsg = err.Error()
			logger.Error("validating request", slog.String("error", errorMsg))
			return c.JSON(http.StatusBadRequest, model.Error{
				Msg: errorMsg,
			})
		}
		token, err := auth.UserToken(req)
		if err != nil {
			errorMsg = err.Error()
			logger.Error("creating token", slog.String("error", errorMsg))
			return c.JSON(http.StatusInternalServerError, model.Error{
				Msg: errorMsg,
			})
		}
		var resp model.LoginResponse

		resp.Token = token

		return c.JSON(http.StatusOK, resp)
	}
}
