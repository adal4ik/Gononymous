package db

import (
	"context"
	"database/sql"

	"backend/internal/core/domains/dao"
)

type SessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepo {
	return &SessionRepo{db: db}
}

func (s *SessionRepo) AddSession(ctx context.Context, session dao.Session) error {
	query := `INSERT INTO users(user_id, name, avatar_url)
			VALUES($1,$2,$3)`
	_, err := s.db.ExecContext(ctx, query, session.UsersId, session.Name, session.AvatarURL)
	if err != nil {
		return err
	}
	return nil
}
