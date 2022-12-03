package main

import (
	"github.com/ferdvtn/go-jwt-auth/configs"
	"github.com/ferdvtn/go-jwt-auth/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	v1 := e.Group("api/v1")

	// controllers
	authCtrl := controllers.NewAuthCtrl()
	userCtrl := controllers.NewUserCtrl()

	// middlewares
	jwt := middleware.JWTWithConfig(configs.JwtConfig)

	// auth
	v1.POST("/login", authCtrl.Login)

	// users
	v1user := v1.Group("/users", jwt)
	v1user.GET("/greetings", userCtrl.Greetings)

	e.Start(":1323")
}
