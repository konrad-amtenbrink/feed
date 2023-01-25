package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a API) Status() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}
}
