package handlers

import (
	chetamApiv1 "chetam/internal/server/client/v1"
	"chetam/internal/validation"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (sh *ServerHandler) PostApiV1AuthRegister(c echo.Context) error {
	var req chetamApiv1.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := validation.ValidateAuthRequest(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//user, err := auth.User(req.Email, req.Login, req.Password)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, err)
	//}
	return nil
}

func (sh *ServerHandler) PostApiV1AuthLogin(c echo.Context) error {
	return c.JSON(http.StatusBadGateway, "fsddsfs")
}
