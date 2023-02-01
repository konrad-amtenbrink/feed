package api

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func (a API) ShowHome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "upload.tmpl", nil)
	}

}

func (a API) ShowReader() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "view.tmpl", nil)
	}
}

func (a API) ShowAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin.tmpl", nil)
	}
}

func (a API) ShowGrid() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "grid.tmpl", nil)
	}
}

func (a API) ShowOverview() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "overview.tmpl", nil)
	}
}
