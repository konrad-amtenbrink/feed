package api

import (
	"html/template"

	"github.com/konrad-amtenbrink/feed/auth"
	"github.com/konrad-amtenbrink/feed/db"
	"github.com/konrad-amtenbrink/feed/storage"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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

	registerFrontend(e, api)

	e.POST("/register", api.Register())
	e.POST("/login", api.Login())

	v1 := e.Group("/v0.1")
	applyAuthMiddleware(v1)

	v1.GET("/documents", api.GetDocuments())
	v1.POST("/documents", api.CreateDocument())

	v1.GET("/document", api.GetDocument())
	v1.DELETE("/document", api.DeleteDocument())

	v1.GET("/status", api.Status())
}

func registerFrontend(e *echo.Echo, api API) {
	e.GET("/", api.ShowOverview())
	e.GET("/:id", api.ShowReader())
	e.GET("/admin", api.ShowAdmin())
	e.GET("/articles", api.ShowGrid())
	e.GET("/create", api.ShowHome())
	e.GET("/login", api.ShowLogin())
	e.GET("/register", api.ShowRegister())
}

func applyAuthMiddleware(e *echo.Group) {
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
