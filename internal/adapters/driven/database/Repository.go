package db

import (
	"database/sql"

	drivenports "Gononymous/internal/core/ports/driven_ports"
)

type Repository struct {
	PostRepo    drivenports.PostDrivenPortInterface
	SessionRepo drivenports.SessionRepoInterface
}

func New(db *sql.DB) *Repository {
	return &Repository{
		PostRepo: NewPostRepository(db),
		// SessionRepo: NewSessionRepo(db),
	}
}
