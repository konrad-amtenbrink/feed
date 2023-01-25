package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	api "github.com/amruid/go-template/api/v0.1"
	"github.com/amruid/go-template/db"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func New(db db.Database) *Server {
	e := echo.New()
	e.HideBanner = true
	e.Validator = &CustomValidator{Validator: validator.New()}

	api.SetupV0_1(e, db)

	return &Server{
		Echo: e,
	}
}

func (srv *Server) Run(port int) error {
	return srv.Echo.Start(":" + strconv.Itoa(port))
}

func (srv *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Echo.Shutdown(ctx); err != nil {
		if closeErr := srv.Echo.Close(); closeErr != nil {
			return fmt.Errorf("shutting down %v and closing server: %v", err, closeErr)
		}

		return fmt.Errorf("shutting down server: %v", err)
	}

	return nil
}
