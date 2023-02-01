package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (db database) CreateUser(ctx context.Context, user User) (uuid.UUID, error) {
	var lastInsertId uuid.UUID
	err := db.Sqlx.GetContext(
		ctx,
		&lastInsertId,
		"Insert into users (username, password, role) values ($1, $2, $3) RETURNING id",
		user.Username, user.Password, user.Role,
	)

	return lastInsertId, err
}

func (db database) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	var user User
	err := db.Sqlx.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", id)
	return user, err
}

func (db database) GetUserByUsername(ctx context.Context, username string) (User, error) {
	var user User
	err := db.Sqlx.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1", username)
	return user, err
}

func (db database) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	_, err := db.Sqlx.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
