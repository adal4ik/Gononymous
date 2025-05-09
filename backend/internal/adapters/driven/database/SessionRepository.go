package db

import (
	"backend/internal/core/domains/dao"
	"context"
	"database/sql"
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

func (s *SessionRepo) GetSessionById(id string) (dao.Session, error) {
	query := `SELECT user_id, name, avatar_url, created_at FROM users WHERE user_id = $1`
	var session dao.Session
	err := s.db.QueryRow(query, id).Scan(&session.UsersId, &session.Name, &session.AvatarURL, &session.CreatedAt)
	if err != nil {
		return dao.Session{}, err
	}
	return session, nil
}
