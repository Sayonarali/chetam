package handlers

import (
	"chetam/internal/model"
	"chetam/internal/validation"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (sh *ServerHandler) Register(c echo.Context) error {
	var req model.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := validation.ValidateAuthRequest(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err = sh.services.Auth.FindUserByLogin(req.Login)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//user, err := auth.User(req.Email, req.Login, req.Password)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, err)
	//}
	return nil
}

func (sh *ServerHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusBadGateway, "fsddsfs")
}
