package test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/adapters/driver/WebHttp/middleware"
)

// Мок реализации SessionServiceDriverInterface
type MockSessionService struct {
	CreateSessionFunc func(ctx context.Context) (string, error)
}

func (m *MockSessionService) CreateSession(ctx context.Context) (string, error) {
	if m.CreateSessionFunc != nil {
		return m.CreateSessionFunc(ctx)
	}
	return "mock-session-id", nil
}

func TestSessionHandler(t *testing.T) {
	t.Run("Session cookie exists", func(t *testing.T) {
		mockSessionService := &MockSessionService{}
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "session_id", Value: "existing-session"})
		res := httptest.NewRecorder()

		middleware := middleware.SessionHandler(handler, mockSessionService)
		middleware.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", res.Code)
		}
	})

	t.Run("Session cookie not present, created successfully", func(t *testing.T) {
		mockSessionService := &MockSessionService{
			CreateSessionFunc: func(ctx context.Context) (string, error) {
				return "new-session-id", nil
			},
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		middleware := middleware.SessionHandler(handler, mockSessionService)
		middleware.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", res.Code)
		}

		cookies := res.Result().Cookies()
		found := false
		for _, c := range cookies {
			if c.Name == "session_id" && c.Value == "new-session-id" {
				found = true
			}
		}
		if !found {
			t.Error("expected session_id cookie to be set")
		}
	})

	t.Run("Session creation error", func(t *testing.T) {
		mockSessionService := &MockSessionService{
			CreateSessionFunc: func(ctx context.Context) (string, error) {
				return "", errors.New("error creating session")
			},
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("handler should not be called on session error")
		})

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		middleware := middleware.SessionHandler(handler, mockSessionService)
		middleware.ServeHTTP(res, req)

		if res.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500, got %d", res.Code)
		}
	})
}
