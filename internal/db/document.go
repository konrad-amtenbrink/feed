package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	UserId    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (db database) CreateDocument(ctx context.Context, doc Document, userId uuid.UUID) (uuid.UUID, error) {
	var lastInsertId uuid.UUID
	err := db.Sqlx.GetContext(
		ctx,
		&lastInsertId,
		"Insert into documents (title, user_id) values ($1, $2) RETURNING id",
		doc.Title, userId,
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

func (db database) GetDocumentsByUserId(ctx context.Context, userId uuid.UUID) ([]Document, error) {
	var docs []Document
	err := db.Sqlx.Select(&docs, "SELECT * FROM documents WHERE user_id = $1 ORDER BY created_at DESC", userId)
	return docs, err
}

func (db database) GetDocuments(ctx context.Context) ([]Document, error) {
	var docs []Document
	err := db.Sqlx.Select(&docs, "SELECT * FROM documents ORDER BY created_at DESC")
	return docs, err
}
