package externalapi

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"backend/internal/core/domains/dto"
)

// CharacterClient — внешний клиент для Rick and Morty API
type CharacterClient struct {
	httpClient *http.Client
}

// NewCharacterClient возвращает новый клиент с отключенной TLS-проверкой (для тестов / dev)
func NewCharacterClient() *CharacterClient {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return &CharacterClient{httpClient: client}
}

// GetCharacter получает персонажа по ID через публичный API
func (c *CharacterClient) GetCharacter(ctx context.Context, id int) (*dto.Character, error) {
	fmt.Println("get character")

	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/%d", id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var character dto.Character
	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		return nil, fmt.Errorf("failed to decode character: %w", err)
	}

	return &character, nil
}
