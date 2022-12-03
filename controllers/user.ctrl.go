package controllers

import (
	"net/http"

	"github.com/ferdvtn/go-jwt-auth/configs"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserCtrl struct{}

func NewUserCtrl() *UserCtrl {
	return &UserCtrl{}
}

func (ctrl UserCtrl) Greetings(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*configs.JwtCustomClaims)

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Welcome " + claims.Username + "!",
	})
}
