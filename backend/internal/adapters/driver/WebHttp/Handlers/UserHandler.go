package handlers

import (
	"backend/internal/core/domains/dao"
	"backend/internal/core/domains/dto"
	driverports "backend/internal/core/ports/driver_ports"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type UserHandler struct {
	service     driverports.UserDriverPortInterface
	session     driverports.SessionServiceDriverInterface
	postService driverports.PostDriverPortInterface
	BaseHandler
}

func NewUserHandler(service driverports.UserDriverPortInterface, baseHandler BaseHandler, session driverports.SessionServiceDriverInterface, postService driverports.PostDriverPortInterface) *UserHandler {
	return &UserHandler{
		service:     service,
		session:     session,
		postService: postService,
		BaseHandler: baseHandler,
	}
}

type ProfilePage struct {
	UserData dao.Session
	Posts    []dto.PostDto
}

type NameChangeRequest struct {
	Name string `json:"name"`
}

func (u *UserHandler) ProfilePage(w http.ResponseWriter, r *http.Request) {
	var page ProfilePage
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Session cookie missing", http.StatusUnauthorized)
		return
	}
	userId := cookie.Value
	page.UserData, err = u.session.GetSessionById(userId)
	if err != nil {
		http.Error(w, "Session cookie missing", http.StatusUnauthorized)
		return
	}
	page.Posts, err = u.postService.GetPostsByUserID(userId)
	if err != nil {
		http.Error(w, "Session cookie missing", http.StatusUnauthorized)
		return
	}
	tmpl := template.Must(template.ParseFiles("web/templates/profile.html"))
	HTMLXposts := map[string]ProfilePage{
		"PostPage": page,
	}
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (u *UserHandler) ChangeName(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Session cookie missing", http.StatusUnauthorized)
		return
	}
	userId := cookie.Value
	var request NameChangeRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newName := strings.TrimSpace(request.Name)
	if newName == "" {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}
	fmt.Println("Here2")
	if len(newName) > 50 {
		http.Error(w, "Name too long (max 50 characters)", http.StatusBadRequest)
		return
	}
	fmt.Println("Here3")
	fmt.Println(userId, " ", newName)
	u.service.ChangeName(userId, newName, r.Context())

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"newName": newName,
	})
}
