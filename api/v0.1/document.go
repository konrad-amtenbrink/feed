package api

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/konrad-amtenbrink/feed/internal/db"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type (
	CreateDocumentResponse struct {
		DocumentId uuid.UUID `json:"document_id"`
	}

	SingleDocumentRequest struct {
		ID string `query:"id"`
	}
)

func (a API) CreateDocument() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser := c.Get("user").(jwt.MapClaims)
		if currentUser["id"] == nil {
			log.Debug("failed to get id from user")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		id, err := uuid.Parse(currentUser["id"].(string))
		if err != nil {
			log.WithError(err).Debug("failed to parse id")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		title := c.FormValue("title")

		file, err := c.FormFile("file")
		if err != nil {
			log.WithError(err).Debug("failed to retrieve file")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		src, err := file.Open()
		if err != nil {
			log.WithError(err).Debug("failed to open file")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		defer src.Close()

		document := db.Document{
			Title: title,
		}

		documentId, err := a.db.CreateDocument(c.Request().Context(), document, id)
		if err != nil {
			log.WithError(err).Debug("failed to create document")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = a.storage.Upload(documentId.String(), src)
		if err != nil {
			log.WithError(err).Debug("failed to create document")

			err = a.db.DeleteDocumentById(c.Request().Context(), documentId)
			if err != nil {
				log.WithError(err).Debug("failed to delete document after creation failed - manual cleanup required")
			}

			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, CreateDocumentResponse{DocumentId: documentId})
	}
}

func (a API) GetDocuments() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser := c.Get("user").(jwt.MapClaims)
		if currentUser["id"] == nil {
			log.Debug("failed to get id from user")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		id, err := uuid.Parse(currentUser["id"].(string))
		if err != nil {
			log.WithError(err).Debug("failed to parse id")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		documents, err := a.db.GetDocumentsByUserId(c.Request().Context(), id)
		if err != nil {
			log.WithError(err).Debug("failed to get documents")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, documents)
	}
}

func (a API) GetAllDocuments() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser := c.Get("user").(jwt.MapClaims)

		if currentUser["role"] != "admin" {
			log.Debug("user is not admin")
			return echo.NewHTTPError(http.StatusForbidden)
		}

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
		var req SingleDocumentRequest
		if err := c.Bind(&req); err != nil {
			log.WithError(err).Debug("failed to bind request")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		id, err := uuid.Parse(req.ID)
		if err != nil {
			log.WithError(err).Debug("failed to parse id")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		document, err := a.db.GetDocumentById(c.Request().Context(), id)
		if err != nil {
			log.WithError(err).Debug("failed to get document")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		data, err := a.storage.Download(document.ID.String())
		if err != nil {
			log.WithError(err).Debug("failed to get document")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.Blob(http.StatusOK, "text/markdown", data)
	}
}

func (a API) DeleteDocument() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req SingleDocumentRequest
		if err := c.Bind(&req); err != nil {
			log.WithError(err).Debug("failed to bind request")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		currentUser := c.Get("user").(jwt.MapClaims)

		if currentUser["role"] != "admin" {
			log.Debug("user is not admin")
			return echo.NewHTTPError(http.StatusForbidden)
		}

		id, err := uuid.Parse(req.ID)
		if err != nil {
			log.WithError(err).Debug("failed to parse id")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		err = a.db.DeleteDocumentById(c.Request().Context(), id)
		if err != nil {
			log.WithError(err).Debug("failed to delete document")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = a.storage.Delete(id.String())
		if err != nil {
			log.WithError(err).Debug("failed to get document")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
