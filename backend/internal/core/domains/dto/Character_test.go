package dto

import (
	"encoding/json"
	"testing"
)

func TestCharacter_MarshalJSON(t *testing.T) {
	// Arrange
	character := &Character{
		Name:      "Rick Sanchez",
		AvatarURL: "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
	}

	// Act
	data, err := json.Marshal(character)
	// Assert
	if err != nil {
		t.Fatalf("Expected no error during marshaling, got: %v", err)
	}

	expected := `{"name":"Rick Sanchez","image":"https://rickandmortyapi.com/api/character/avatar/1.jpeg"}`
	if string(data) != expected {
		t.Fatalf("Expected JSON: %s, got: %s", expected, data)
	}
}

func TestCharacter_UnmarshalJSON(t *testing.T) {
	// Arrange
	jsonData := `{"name":"Rick Sanchez","image":"https://rickandmortyapi.com/api/character/avatar/1.jpeg"}`

	// Act
	var character Character
	err := json.Unmarshal([]byte(jsonData), &character)
	// Assert
	if err != nil {
		t.Fatalf("Expected no error during unmarshaling, got: %v", err)
	}

	if character.Name != "Rick Sanchez" {
		t.Errorf("Expected name: 'Rick Sanchez', got: %s", character.Name)
	}
	if character.AvatarURL != "https://rickandmortyapi.com/api/character/avatar/1.jpeg" {
		t.Errorf("Expected AvatarURL: 'https://rickandmortyapi.com/api/character/avatar/1.jpeg', got: %s", character.AvatarURL)
	}
}

func TestCharacter_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	jsonData := `{"name":"Rick Sanchez"}` // Missing image URL

	// Act
	var character Character
	err := json.Unmarshal([]byte(jsonData), &character)
	// Assert
	if err != nil {
		t.Fatalf("Expected no error during unmarshaling, got: %v", err)
	}

	// Ensure AvatarURL is empty since it's missing in the input JSON
	if character.AvatarURL != "" {
		t.Errorf("Expected AvatarURL to be empty, got: %s", character.AvatarURL)
	}
}
