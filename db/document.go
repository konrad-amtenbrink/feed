package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (db database) CreateDocument(ctx context.Context, doc Document) (uuid.UUID, error) {
	var lastInsertId uuid.UUID
	err := db.Sqlx.GetContext(
		ctx,
		&lastInsertId,
		"Insert into documents (title) values ($1) RETURNING id",
		doc.Title,
	)

	return lastInsertId, err
}

func (db database) GetDocumentById(ctx context.Context, id uuid.UUID) (Document, error) {
	var doc Document
	err := db.Sqlx.GetContext(ctx, &doc, "SELECT * FROM documents WHERE id = $1", id)
	return doc, err
}

func (db database) DeleteDocumentById(ctx context.Context, id uuid.UUID) error {
	_, err := db.Sqlx.ExecContext(ctx, "DELETE FROM documents WHERE id = $1", id)
	return err
}

func (db database) GetDocuments(ctx context.Context) ([]Document, error) {
	var docs []Document
	err := db.Sqlx.Select(&docs, "SELECT * FROM documents ORDER BY created_at DESC")
	return docs, err
}
