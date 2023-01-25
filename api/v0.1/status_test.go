package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/amruid/go-template/db"
	"github.com/golang/mock/gomock"
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

	api := API{db}

	resp := api.Status()(ctx)
	if assert.NoError(t, resp) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "ok", rec.Body.String())
	}
}
