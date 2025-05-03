package handlers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"backend/internal/core/domains/dto"

	driverports "backend/internal/core/ports/driver_ports"
)

type PostsHandler struct {
	service driverports.PostDriverPortInterface
	BaseHandler
}

func NewPostHandler(service driverports.PostDriverPortInterface, baseHandler BaseHandler) *PostsHandler {
	return &PostsHandler{service: service, BaseHandler: baseHandler}
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
	var post dto.PostDto
	post.Title = r.Form["name"][0]
	post.Subject = r.Form["subject"][0]
	post.Content = r.Form["comment"][0]
	in, _, err := r.FormFile("file")
	defer in.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	cookie, err := r.Cookie("session_id")
	post.AuthorID = cookie.Value
	data, err := io.ReadAll(in)
	err = postHandler.service.AddPost(post, data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (postsHandler *PostsHandler) PostPage(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")
	post, err := postsHandler.service.GetPostById(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl := template.Must(template.ParseFiles("web/templates/post.html"))
	HTMLXposts := map[string]dto.PostDto{
		"Post": post,
	}
	fmt.Println(HTMLXposts)
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
