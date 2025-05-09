package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"backend/internal/core/domains/dto"
	driverports "backend/internal/core/ports/driver_ports"
)

type ArchiveHandler struct {
	service driverports.PostDriverPortInterface
	comment driverports.CommentServiceInterface
	session driverports.SessionServiceDriverInterface
	BaseHandler
}

func NewArchiveHandler(service driverports.PostDriverPortInterface, baseHandler BaseHandler, comment driverports.CommentServiceInterface, session driverports.SessionServiceDriverInterface) *ArchiveHandler {
	return &ArchiveHandler{
		service:     service,
		BaseHandler: baseHandler,
		comment:     comment,
		session:     session,
	}
}

func (h *ArchiveHandler) GetArchivePage(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetAll()
	if err != nil {
		h.handleError(w, r, 500, "Failed to get posts", err)
		h.RenderError(w, 500, "Failed to get posts")
		return
	}
	tmpl := template.Must(template.ParseFiles("web/templates/archive.html"))
	HTMLXposts := map[string][]dto.PostDto{
		"Posts": posts,
	}
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		h.handleError(w, r, 500, "Failed to Execute", err)
		h.RenderError(w, 500, "Failes to Execute")
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err != nil {
		h.handleError(w, r, 500, "Failed to write response", err)
		h.RenderError(w, 500, "Failed to write response")
		return
	}
	h.logger.Info("Archive page rendered successfully", "url", r.URL.Path)
}

func (h *ArchiveHandler) GetArchivePost(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")
	var page PostPage
	var err error
	page.Post, err = h.service.GetPostById(postId)
	if err != nil {
		h.handleError(w, r, 500, "Failed to get post", err)
		h.RenderError(w, 500, "Failed to get post")
		return
	}
	page.User, err = h.session.GetSessionById(page.Post.AuthorID)
	if err != nil {
		h.handleError(w, r, 500, "Failed to get user", err)
		h.RenderError(w, 500, "Failde to get user")
		return
	}
	page.Comments, err = h.comment.GetCommentsByPostId(postId)
	if err != nil {
		h.handleError(w, r, 500, "Failed to get comment", err)
		h.RenderError(w, 500, "Failed to get comment")
		return
	}
	tmpl := template.Must(template.ParseFiles("web/templates/archive-post.html"))
	HTMLXposts := map[string]PostPage{
		"PostPage": page,
	}
	fmt.Println(HTMLXposts)
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		h.handleError(w, r, 500, "Failed to get comment", err)
		h.RenderError(w, 500, "Failed to get comment")
	}
}
