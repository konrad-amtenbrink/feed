package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	URL       string    `json:"url" db:"url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (db database) CreateDocument(ctx context.Context, doc Document) (uuid.UUID, error) {
	var lastInsertId uuid.UUID
	err := db.Sqlx.GetContext(
		ctx,
		&lastInsertId,
		"Insert into documents (title, url) values ($1, $2) RETURNING id",
		doc.Title, doc.URL,
	)

	return lastInsertId, err
}

func (db database) GetDocumentById(ctx context.Context, id uuid.UUID) (Document, error) {
	var doc Document
	err := db.Sqlx.GetContext(ctx, &doc, "SELECT * FROM documents WHERE id = $1", id)
	return doc, err
}
