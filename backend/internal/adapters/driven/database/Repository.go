package db

import (
	"database/sql"

	externalapi "backend/internal/adapters/driven/external_api"
	drivenports "backend/internal/core/ports/driven_ports"
)

type Repository struct {
	PostRepo      drivenports.DatabasePortInterface
	SessionRepo   drivenports.SessionRepoInterface
	CharacterRepo drivenports.CharacterRepoInterface
	CommentRepo   drivenports.CommentRepoInterface
}

func New(db *sql.DB) *Repository {
	return &Repository{
		PostRepo:      NewPostRepository(db),
		SessionRepo:   NewSessionRepo(db),
		CharacterRepo: externalapi.NewCharacterClient(),
		CommentRepo:   NewCommentRepository(db),
	}
}
