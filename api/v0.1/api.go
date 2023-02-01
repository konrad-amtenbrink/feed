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

	e.GET("/", api.ShowOverview())
	e.GET("/:id", api.ShowReader())
	e.GET("/admin", api.ShowAdmin())
	e.GET("/articles", api.ShowGrid())
	e.GET("/create", api.ShowHome())

	v1 := e.Group("/v0.1")

	v1.GET("/documents", api.GetDocuments())
	v1.POST("/documents", api.CreateDocument())

	v1.GET("/document", api.GetDocument())
	v1.DELETE("/document", api.DeleteDocument())

	v1.POST("/register", api.Register())
	v1.POST("/login", api.Login())

	v1.GET("/status", api.Status())
}
