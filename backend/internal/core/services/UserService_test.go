package services

import (
	"context"
	"errors"
	"testing"
)

// Mock for UserDriverPortInterface
type MockUserRepo struct {
	ChangeNameFn func(userId string, newName string, ctx context.Context) error
}

func (m *MockUserRepo) ChangeName(userId string, newName string, ctx context.Context) error {
	if m.ChangeNameFn != nil {
		return m.ChangeNameFn(userId, newName, ctx)
	}
	return nil
}

func TestChangeName_Success(t *testing.T) {
	// Arrange
	mockUserRepo := &MockUserRepo{
		ChangeNameFn: func(userId string, newName string, ctx context.Context) error {
			if userId == "" || newName == "" {
				return errors.New("invalid parameters")
			}
			return nil
		},
	}
	userService := NewUserService(mockUserRepo)

	// Act
	err := userService.ChangeName("123", "NewName", context.Background())
	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestChangeName_Error(t *testing.T) {
	// Arrange
	mockUserRepo := &MockUserRepo{
		ChangeNameFn: func(userId string, newName string, ctx context.Context) error {
			if userId == "" || newName == "" {
				return errors.New("invalid parameters")
			}
			return nil
		},
	}
	userService := NewUserService(mockUserRepo)

	// Act
	err := userService.ChangeName("", "", context.Background())

	// Assert
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if err.Error() != "invalid parameters" {
		t.Fatalf("Expected 'invalid parameters' error, got %v", err)
	}
}
