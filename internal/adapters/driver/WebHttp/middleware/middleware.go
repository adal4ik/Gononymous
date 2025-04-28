package middleware

import (
	"Gononymous/utils"
	"fmt"
	"net/http"
	"time"
)

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		sessionID := utils.UUID()
		newCookie := &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		}
		http.SetCookie(w, newCookie)
		fmt.Println(newCookie)
		return
	}
	fmt.Println(cookie)

	return
}
