package handlers

import (
	"backend/internal/core/domains/dao"
	"backend/internal/core/domains/dto"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	driverports "backend/internal/core/ports/driver_ports"
)

type PostsHandler struct {
	service  driverports.PostDriverPortInterface
	comments driverports.CommentServiceInterface
	session  driverports.SessionServiceDriverInterface
	BaseHandler
}

func NewPostHandler(service driverports.PostDriverPortInterface, comments driverports.CommentServiceInterface, session driverports.SessionServiceDriverInterface, baseHandler BaseHandler) *PostsHandler {
	return &PostsHandler{service: service, comments: comments, session: session, BaseHandler: baseHandler}
}

type PostPage struct {
	User     dao.Session
	Post     dto.PostDto
	Comments []dto.Comment
}

func (postHandler *PostsHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/create-post.html"))
	tmpl.Execute(w, nil)
}

func (postHandler *PostsHandler) SubmitPostHandler(w http.ResponseWriter, r *http.Request) {
	var size int64
	size = r.ContentLength
	if err := r.ParseMultipartForm(size); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var err error
	var img []byte
	var post dto.PostDto

	post.Title = strings.TrimSpace(r.FormValue("name"))
	if post.Title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	post.Subject = strings.TrimSpace(r.FormValue("subject"))
	post.Content = strings.TrimSpace(r.FormValue("comment"))

	file, header, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		if header.Size > 10<<20 {
			http.Error(w, "File too large (max 10MB)", http.StatusBadRequest)
			return
		}

		img, err = io.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
	} else if err != http.ErrMissingFile {
		http.Error(w, "Error processing file upload", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Session cookie missing", http.StatusUnauthorized)
		return
	}
	post.AuthorID = cookie.Value

	err = postHandler.service.AddPost(post, img)
	if err != nil {
		http.Error(w, "Error saving post", http.StatusInternalServerError)
		return
	}
}

func (postsHandler *PostsHandler) PostPage(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")
	var page PostPage
	var err error
	page.Post, err = postsHandler.service.GetPostById(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	page.User, err = postsHandler.session.GetSessionById(page.Post.AuthorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	page.Comments, err = postsHandler.comments.GetCommentsByPostId(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl := template.Must(template.ParseFiles("web/templates/post.html"))
	HTMLXposts := map[string]PostPage{
		"PostPage": page,
	}
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
