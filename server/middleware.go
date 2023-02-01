package server

import (
	"github.com/konrad-amtenbrink/feed/auth"
	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApplyMiddleware(e *echo.Echo) {
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			access_token, err := c.Cookie("access_token")
			if err != nil {
				log.WithError(err).Debug("failed to get access_token")
				return err
			}

			user, err := auth.Parse(c, access_token.Value)
			if err != nil {
				log.WithError(err).Debug("failed to parse access_token")
				return err
			}
			c.Set("user", user)
			return next(c)

		}
	})
}
