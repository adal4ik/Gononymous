package handlers

import (
	"html/template"
	"io"
	"net/http"

	"backend/internal/core/domains/dao"
	"backend/internal/core/domains/dto"
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
	post.Title = r.Form["name"][0]
	post.Subject = r.Form["subject"][0]
	post.Content = r.Form["comment"][0]
	if len(r.FormValue("file")) != 0 {
		in, _, err := r.FormFile("file")
		if err != nil {
			postHandler.handleError(w, r, 500, "asd", err)
			postHandler.RenderError(w, 500, "asd")
			return
		}
		defer in.Close()
		img, err = io.ReadAll(in)
	}
	if err != nil {
		postHandler.handleError(w, r, 500, "asd", err)
		postHandler.RenderError(w, 500, "asd")
		return
	}
	cookie, err := r.Cookie("session_id")
	post.AuthorID = cookie.Value

	err = postHandler.service.AddPost(post, img)
	if err != nil {
		postHandler.handleError(w, r, 500, "asd", err)
		postHandler.RenderError(w, 500, "asd")
		return
	}
}

func (postsHandler *PostsHandler) PostPage(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")
	var page PostPage
	var err error
	page.Post, err = postsHandler.service.GetPostById(postId)
	if err != nil {
		postsHandler.handleError(w, r, 500, "asd", err)
		postsHandler.RenderError(w, 500, "asd")
		return
	}
	page.User, err = postsHandler.session.GetSessionById(page.Post.AuthorID)
	if err != nil {
		postsHandler.handleError(w, r, 500, "asd", err)
		postsHandler.RenderError(w, 500, "asd")
		return
	}
	page.Comments, err = postsHandler.comments.GetCommentsByPostId(postId)
	if err != nil {
		postsHandler.handleError(w, r, 500, "asd", err)
		postsHandler.RenderError(w, 500, "asd")
		return
	}
	tmpl := template.Must(template.ParseFiles("web/templates/post.html"))
	HTMLXposts := map[string]PostPage{
		"PostPage": page,
	}
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		postsHandler.handleError(w, r, 500, "asd", err)
		postsHandler.RenderError(w, 500, "asd")
		return
	}
}
