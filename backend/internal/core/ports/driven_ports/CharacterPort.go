package drivenports

import (
	"context"

	"backend/internal/core/domains/dto"
)

type CharacterRepoInterface interface {
	GetCharacter(ctx context.Context, id int) (*dto.Character, error)
}
