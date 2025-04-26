package db

import (
	"database/sql"

	drivenports "Gononymous/internal/core/ports/driven_ports"
)

type Repository struct {
	PostRepo drivenports.PostDrivenPortInterface
}

func New(db *sql.DB) *Repository {
	return &Repository{
		PostRepo: NewPostRepository(db),
	}
}
