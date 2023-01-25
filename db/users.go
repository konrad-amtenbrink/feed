package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Firstname string    `json:"firstname" db:"firstname"`
	Lastname  string    `json:"lastname" db:"lastname"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (db database) CreateUser(ctx context.Context, user User) (uuid.UUID, error) {
	var lastInsertId uuid.UUID
	err := db.Sqlx.GetContext(
		ctx,
		&lastInsertId,
		"Insert into users (firstname, lastname, email) values ($1, $2, $3) RETURNING id",
		user.Firstname, user.Lastname, user.Email,
	)

	return lastInsertId, err
}

func (db database) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	var user User
	err := db.Sqlx.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", id)
	return user, err
}
