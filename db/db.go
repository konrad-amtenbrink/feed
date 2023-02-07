//go:generate mockgen -source=$GOFILE -package=db -destination=db_mock.go
package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/konrad-amtenbrink/feed/config"
	_ "github.com/lib/pq"
)

type (
	Database interface {
		CreateDocument(ctx context.Context, doc Document, userId uuid.UUID) (uuid.UUID, error)
		GetDocumentsByUserId(ctx context.Context, userId uuid.UUID) ([]Document, error)
		DeleteDocumentById(ctx context.Context, id uuid.UUID) error
		GetDocumentById(ctx context.Context, id uuid.UUID) (Document, error)
		GetUserById(ctx context.Context, id uuid.UUID) (User, error)
		GetUserByUsername(ctx context.Context, username string) (User, error)
		CreateUser(ctx context.Context, user User) (uuid.UUID, error)
		DeleteUserById(ctx context.Context, id uuid.UUID) error
	}

	CloseFunc func() error

	database struct {
		Sqlx *sqlx.DB
	}
)

func NewDatabase(ctx context.Context, cfg config.DBConfig) (Database, CloseFunc, error) {
	db, err := connect(ctx, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("connecting to database: %v", err)
	}

	err = runMigrations(db.DB)
	if err != nil {
		return nil, nil, fmt.Errorf("running migrations: %v", err)
	}

	closeFunc := func() error {
		return db.Close()
	}
	return database{db}, closeFunc, nil
}

func connect(ctx context.Context, cfg config.DBConfig) (*sqlx.DB, error) {
	database, err := sqlx.ConnectContext(ctx, "postgres", newConnectionString(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return database, nil
}

func runMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("create postgres driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("create new migrate with db instance: %v", err)
	}

	
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("running the migration: %v", err)
	}

	return nil
}

func newConnectionString(cfg config.DBConfig) string {
	var sslMode string
	if cfg.DisableSslMode {
		sslMode = "disable"
	} else {
		sslMode = "require"
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.ConnectTimeout,
		sslMode,
	)
}
