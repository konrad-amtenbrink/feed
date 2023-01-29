package api

import (
	"html/template"

	"github.com/konrad-amtenbrink/feed/db"
	"github.com/konrad-amtenbrink/feed/storage"
	"github.com/labstack/echo/v4"
)

type API struct {
	db      db.Database
	storage storage.Storage
}

func SetupV0_1(e *echo.Echo, db db.Database, storage storage.Storage) {
	// see api/v0.1/views.go
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.tmpl")),
	}
	api := API{
		db:      db,
		storage: storage,
	}

	e.Static("/static", "static")
	e.Renderer = renderer

	e.GET("/create", api.ShowHome())
	e.GET("/:id", api.ShowReader())
	e.GET("/admin", api.ShowBrowse())
	e.GET("/", api.ShowGrid())

	v1 := e.Group("/v0.1")

	v1.POST("/documents", api.CreateDocument())
	v1.GET("/document", api.GetDocument())
	v1.GET("/documents", api.GetDocuments())

	v1.GET("/status", api.Status())
}
