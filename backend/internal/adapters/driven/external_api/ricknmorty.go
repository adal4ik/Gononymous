package externalapi

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"backend/internal/core/domains/dto"
)

type RickAndMortySerivice struct {
	client *http.Client
}

func NewRickAndMortySerivice() *RickAndMortySerivice {
	customCLinet := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	return &RickAndMortySerivice{
		client: customCLinet,
	}
}

func (r *RickAndMortySerivice) GetCharacter(ctx context.Context, id int) (*dto.Character, error) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/%d", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var character dto.Character
	if err := json.NewDecoder(res.Body).Decode(&character); err != nil {
		return nil, err
	}
	return &character, nil
}
