package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type (
	SingleUserRequest struct {
		ID string `query:"id"`
	}
)

func (a API) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req SingleUserRequest
		if err := c.Bind(&req); err != nil {
			log.WithError(err).Debug("failed to bind request")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		id, err := uuid.Parse(req.ID)
		if err != nil {
			log.WithError(err).Debug("failed to parse id")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		user, err := a.db.GetUserById(c.Request().Context(), id)
		if err != nil {
			log.WithError(err).Debug("failed to get user")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, user)
	}
}
