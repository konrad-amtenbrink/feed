package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/konrad-amtenbrink/feed/db"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type (
	CreateUserRequest struct {
		Firstname string `json:"firstname" validate:"required"`
		Lastname  string `json:"lastname" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
	}

	CreateUserResponse struct {
		UserId uuid.UUID `json:"user_id"`
	}

	GetUserRequest struct {
		UserId uuid.UUID `param:"id" json:"user_id" validate:"required"`
	}
)

func (a API) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req CreateUserRequest
		if err := c.Bind(&req); err != nil {
			log.WithError(err).Debug("failed to bind request")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if err := c.Validate(req); err != nil {
			log.WithError(err).Debug("failed to validate request")
			return err
		}

		user := db.User{
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
			Email:     req.Email,
		}
		userId, err := a.db.CreateUser(c.Request().Context(), user)
		if err != nil {
			log.WithError(err).Debug("failed to create user")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, CreateUserResponse{UserId: userId})
	}
}

func (a API) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GetUserRequest
		if err := c.Bind(&req); err != nil {
			log.WithError(err).Debug("failed to bind request")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if err := c.Validate(req); err != nil {
			log.WithError(err).Debug("failed to validate request")
			return err
		}

		user, err := a.db.GetUserById(c.Request().Context(), req.UserId)
		if err != nil {
			log.WithError(err).Debug("failed to get user")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, user)
	}
}
