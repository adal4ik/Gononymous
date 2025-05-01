package drivenports

import (
	"Gononymous/internal/core/domains/dto"
	"context"
)

type CharacterRepoInterface interface {
	GetCharacter(ctx context.Context, id int) (*dto.Character, error)
}
