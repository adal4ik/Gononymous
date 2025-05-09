package handlers

import (
	"backend/internal/core/domains/dao"
	"backend/internal/core/domains/dto"
	driverports "backend/internal/core/ports/driver_ports"
	"encoding/json"
	"fmt"
	"html/template"
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
		u.handleError(w, r, 500, "Fail", err)
		u.RenderError(w, 500, "Fail")
		return
	}
	userId := cookie.Value
	page.UserData, err = u.session.GetSessionById(userId, r.Context())
	if err != nil {
		u.handleError(w, r, 401, "Fail", err)
		u.RenderError(w, 401, "Fail")
		return
	}
	page.Posts, err = u.postService.GetPostsByUserID(userId, r.Context())
	if err != nil {
		u.handleError(w, r, 401, "Fail", err)
		u.RenderError(w, 401, "Fail")
		return
	}
	tmpl := template.Must(template.ParseFiles("web/templates/profile.html"))
	HTMLXposts := map[string]ProfilePage{
		"PostPage": page,
	}
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		u.handleError(w, r, 500, "Fail", err)
		u.RenderError(w, 500, "Fail")
	}
}

func (u *UserHandler) ChangeName(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		u.handleError(w, r, 401, "Fail", err)
		u.RenderError(w, 401, "Fail")
		return
	}
	userId := cookie.Value
	var request NameChangeRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		u.handleError(w, r, 400, "Fail", err)
		u.RenderError(w, 400, "Fail")
		return
	}
	newName := strings.TrimSpace(request.Name)
	if newName == "" {
		u.handleError(w, r, 400, "Fail", err)
		u.RenderError(w, 400, "Fail")
		return
	}

	if len(newName) > 50 {
		u.handleError(w, r, 400, "Fail <50", err)
		u.RenderError(w, 400, "Fail <50")
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
