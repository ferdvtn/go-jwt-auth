package controllers

import (
	"net/http"
	"time"

	"github.com/ferdvtn/go-jwt-auth/configs"
	"github.com/ferdvtn/go-jwt-auth/domains"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthCtrl struct{}

func NewAuthCtrl() *AuthCtrl {
	return &AuthCtrl{}
}

func (auth AuthCtrl) Login(ctx echo.Context) error {
	u := ctx.FormValue("username")
	p := ctx.FormValue("password")

	if u != "fer" || p != "123" {
		return echo.ErrUnauthorized
	}

	user := domains.User{Id: 1, Username: u, IsAdmin: true, Name: "Ahmad Ferdian"}

	// Set custom claims
	claims := configs.JwtCustomClaims{
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 72),
			},
		},
	}

	// Create token w/ claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	t, err := token.SignedString(configs.JwtKey)
	if err != nil {
		return echo.ErrUnauthorized
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": t,
		"user":  user,
	})
}
