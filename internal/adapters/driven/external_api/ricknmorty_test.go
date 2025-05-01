package externalapi

import (
	"context"
	"testing"
)

func TestRickAndMortySerivice_GetCharacter(t *testing.T) {
	ctx := context.Background()
	r := NewRickAndMortySerivice()

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
