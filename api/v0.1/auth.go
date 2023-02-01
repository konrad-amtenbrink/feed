package api

import (
	"net/http"

	"github.com/konrad-amtenbrink/feed/auth"
	"github.com/konrad-amtenbrink/feed/db"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type (
	AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func (a API) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req AuthRequest
		if err := c.Bind(&req); err != nil {
			log.WithError(err).Debug("failed to bind request")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		hash, err := auth.HashAndSalt(req.Password)
		if err != nil {
			log.WithError(err).Debug("failed to hash password")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		user, err := a.db.CreateUser(c.Request().Context(), db.User{Username: req.Username, Password: hash, Role: "user"})
		if err != nil {
			log.WithError(err).Debug("failed to get user")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (a API) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req AuthRequest
		if err := c.Bind(&req); err != nil {
			log.WithError(err).Debug("failed to bind request")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		user, err := a.db.GetUserByUsername(c.Request().Context(), req.Username)
		if err != nil {
			log.WithError(err).Debug("failed to get user")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		if !auth.ComparePasswords(user.Password, req.Password) {
			log.WithError(err).Debug("wrong password for user %s", user.Username)
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		err = auth.GenerateAndSet(c, user)
		if err != nil {
			log.WithError(err).Debug("failed to generate token")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, user)
	}
}
