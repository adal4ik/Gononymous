package handlers

import (
	"Gononymous/internal/core/domains/dto"
	"fmt"
	"html/template"
	"net/http"

	driverports "Gononymous/internal/core/ports/driver_ports"
)

type PostsHandler struct {
	service driverports.PostDriverPortInterface
}

func NewPostHandler(service driverports.PostDriverPortInterface) *PostsHandler {
	return &PostsHandler{service: service}
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
	post.Image = r.Form["file"][0]
	err := postHandler.service.AddPost(post)
	if err != nil {
		fmt.Println(err.Error())
	}
}
