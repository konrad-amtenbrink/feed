package server

import (
	"github.com/konrad-amtenbrink/feed/auth"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApplyMiddleware(e *echo.Echo) {
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(echojwt.WithConfig(echojwt.Config{
		ParseTokenFunc: auth.Parse,
	}))
}
