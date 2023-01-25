package api

import (
	"github.com/konrad-amtenbrink/feed/db"
	"github.com/labstack/echo/v4"
)

type API struct {
	db db.Database
}

func SetupV0_1(e *echo.Echo, db db.Database) {
	api := API{
		db: db,
	}

	v1 := e.Group("/v0.1")

	v1.POST("/users", api.CreateUser())
	v1.GET("/users/:id", api.GetUser())

	v1.GET("/status", api.Status())
}
