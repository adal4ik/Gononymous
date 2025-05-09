package handlers

import (
	"backend/internal/core/domains/dao"
	"backend/internal/core/domains/dto"
	"html/template"
	"io"
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
		postHandler.handleError(w, r, 500, "asd", err)
		postHandler.RenderError(w, 500, "asd")
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
			postHandler.handleError(w, r, 500, "to much size of file", err)
			postHandler.RenderError(w, 500, "to much size of file")
			return
		}

		img, err = io.ReadAll(file)
		if err != nil {
			postHandler.handleError(w, r, 500, "Fail", err)
			postHandler.RenderError(w, 500, "Fail")
			return
		}
	} else if err != http.ErrMissingFile {
		postHandler.handleError(w, r, 400, "Fail", err)
		postHandler.RenderError(w, 400, "Fail")
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		postHandler.handleError(w, r, 500, "Fail", err)
		postHandler.RenderError(w, 500, "Fail")
		return
	}
	post.AuthorID = cookie.Value

	err = postHandler.service.AddPost(post, img, r.Context())
	if err != nil {
		postHandler.handleError(w, r, 500, "Fail", err)
		postHandler.RenderError(w, 500, "Fail")
		return
	}
}

func (postsHandler *PostsHandler) PostPage(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")
	var page PostPage
	var err error
	page.Post, err = postsHandler.service.GetPostById(postId, r.Context())
	if err != nil {
		postsHandler.handleError(w, r, 500, "Fail", err)
		postsHandler.RenderError(w, 500, "Fail")
		return
	}
	page.User, err = postsHandler.session.GetSessionById(page.Post.AuthorID, r.Context())
	if err != nil {
		postsHandler.handleError(w, r, 500, "Fail", err)
		postsHandler.RenderError(w, 500, "Fail")
		return
	}
	page.Comments, err = postsHandler.comments.GetCommentsByPostId(postId, r.Context())
	if err != nil {
		postsHandler.handleError(w, r, 500, "Fail", err)
		postsHandler.RenderError(w, 500, "Fail")
		return
	}
	tmpl := template.Must(template.ParseFiles("web/templates/post.html"))
	HTMLXposts := map[string]PostPage{
		"PostPage": page,
	}
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		postsHandler.handleError(w, r, 500, "Fail", err)
		postsHandler.RenderError(w, 500, "Fail")
		return
	}
}
