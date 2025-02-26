package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
)

func GetUser(logger *slog.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
