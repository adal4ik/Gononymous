package handlers

import (
	"backend/internal/core/domains/dao"
	"backend/internal/core/domains/dto"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

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
	post.Title = r.Form["name"][0]
	post.Subject = r.Form["subject"][0]
	post.Content = r.Form["comment"][0]
	if len(r.FormValue("file")) != 0 {
		in, _, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer in.Close()
		img, err = io.ReadAll(in)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cookie, err := r.Cookie("session_id")
	post.AuthorID = cookie.Value

	err = postHandler.service.AddPost(post, img)
	if err != nil {
		fmt.Println(err.Error())
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
