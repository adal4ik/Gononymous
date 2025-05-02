package db

import (
	"database/sql"

	drivenports "backend/internal/core/ports/driven_ports"
)

type Repository struct {
	PostRepo    drivenports.DatabasePortInterface
	SessionRepo drivenports.SessionRepoInterface
}

func New(db *sql.DB) *Repository {
	return &Repository{
		PostRepo: NewPostRepository(db),
		// SessionRepo: NewSessionRepo(db),
	}
}
