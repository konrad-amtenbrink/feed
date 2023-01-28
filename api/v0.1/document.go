package api

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/konrad-amtenbrink/feed/db"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type (
	CreateDocumentResponse struct {
		DocumentId uuid.UUID `json:"document_id"`
	}

	GetDocumentRequest struct {
		ID string `query:"id"`
	}
)

func (a API) CreateDocument() echo.HandlerFunc {
	return func(c echo.Context) error {
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

		documentId, err := a.db.CreateDocument(c.Request().Context(), document)
		if err != nil {
			log.WithError(err).Debug("failed to create document")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		fileType := strings.Split(file.Filename, ".")[1]
		updatedFileId := documentId.String() + "." + fileType

		err = a.storage.Upload(updatedFileId, src)
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
