package configs

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JwtKey = []byte("secret123")

type JwtCustomClaims struct {
	Username string
	IsAdmin  bool
	jwt.RegisteredClaims
}

var JwtConfig = middleware.JWTConfig{
	Claims:     &JwtCustomClaims{},
	SigningKey: JwtKey,
}
