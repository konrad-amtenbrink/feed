package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/konrad-amtenbrink/feed/db"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type (
	CreateDocumentRequest struct {
		Title string `json:"title" validate:"required"`
		URL   string `json:"url" validate:"required"`
	}

	CreateDocumentResponse struct {
		DocumentId uuid.UUID `json:"document_id"`
	}

	GetDocumentRequest struct {
		DocumentId uuid.UUID `param:"id" json:"document_id" validate:"required"`
	}

	GetDocumentResponse struct {
		DocumentTitle   string `json:"document_title"`
		DocumentContent string `json:"document_content"`
	}
)

func (a API) CreateDocument() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req CreateDocumentRequest
		if err := c.Bind(&req); err != nil {
			log.WithError(err).Debug("failed to bind request")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if err := c.Validate(req); err != nil {
			log.WithError(err).Debug("failed to validate request")
			return err
		}

		document := db.Document{
			Title: req.Title,
			URL:   req.URL,
		}
		documentId, err := a.db.CreateDocument(c.Request().Context(), document)
		if err != nil {
			log.WithError(err).Debug("failed to create document")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, CreateDocumentResponse{DocumentId: documentId})
	}
}

func (a API) GetDocuments() echo.HandlerFunc {
	return func(c echo.Context) error {
		documents, err := a.db.GetDocuments(c.Request().Context())
		if err != nil {
			log.WithError(err).Debug("failed to get documents")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, documents)
	}
}

func (a API) GetDocument() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GetDocumentRequest
		if err := c.Bind(&req); err != nil {
			log.WithError(err).Debug("failed to bind request")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if err := c.Validate(req); err != nil {
			log.WithError(err).Debug("failed to validate request")
			return err
		}

		document, err := a.db.GetDocumentById(c.Request().Context(), req.DocumentId)
		if err != nil {
			log.WithError(err).Debug("failed to get document")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, document)
	}
}
