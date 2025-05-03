package middleware

import (
	"net/http"
	"time"

	driverports "backend/internal/core/ports/driver_ports"
)

func SessionHandler(next http.Handler, sessionService driverports.SessionServiceDriverInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session_id")
		if err != nil && err != http.ErrNoCookie {
			http.Error(w, "error reading cookie", http.StatusBadRequest)
			return
		}

		if err == http.ErrNoCookie {

			id, err := sessionService.CreateSession(r.Context())
			if err != nil {
				http.Error(w, "failed to create session", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    id,
				Path:     "/",
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
			})
		}

		next.ServeHTTP(w, r)
	})
}
