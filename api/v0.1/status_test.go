package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/konrad-amtenbrink/feed/internal/db"
	"github.com/konrad-amtenbrink/feed/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Status(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	db := db.NewMockDatabase(ctrl)
	storage := storage.NewMockStorage(ctrl)

	api := API{db, storage}

	resp := api.Status()(ctx)
	if assert.NoError(t, resp) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "ok", rec.Body.String())
	}
}
