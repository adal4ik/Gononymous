package db

import (
	"context"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) ChangeName(userId string, newName string, ctx context.Context) error {
	_, err := u.db.ExecContext(ctx, "UPDATE users SET name = $1 WHERE user_id = $2", newName, userId)
	if err != nil {
		return err
	}
	return nil
}
