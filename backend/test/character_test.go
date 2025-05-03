package test

import (
	"context"
	"testing"

	externalapi "backend/internal/adapters/driven/external_api"
)

func TestCharacterSerivice_GetCharacter(t *testing.T) {
	ctx := context.Background()
	r := externalapi.NewCharacterClient()

	ch, err := r.GetCharacter(ctx, 1)
	if err != nil {
		t.Fatal("Failed to ger charcater: %w", err)
	}

	if ch.Name != "Rick Sanchez" {
		t.Errorf("Expected Rick Sanchez, got %s", ch.Name)
	}

	if ch.AvatarURL != "https://rickandmortyapi.com/api/character/avatar/1.jpeg" {
		t.Error("Wrong Image URL")
	}
	t.Log(ch)
}
