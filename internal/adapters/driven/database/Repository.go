package db

import (
	"database/sql"

	externalapi "Gononymous/internal/adapters/driven/external_api"
	drivenports "Gononymous/internal/core/ports/driven_ports"
)

type Repository struct {
	PostRepo      drivenports.PostDrivenPortInterface
	SessionRepo   drivenports.SessionRepoInterface
	CharacterRepo drivenports.CharacterRepoInterface
}

func New(db *sql.DB) *Repository {
	return &Repository{
		PostRepo:      NewPostRepository(db),
		SessionRepo:   NewSessionRepo(db),
		CharacterRepo: externalapi.NewCharacterClient(),
	}
}
