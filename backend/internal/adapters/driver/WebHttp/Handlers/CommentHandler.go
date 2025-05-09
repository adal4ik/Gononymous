package handlers

import (
	"backend/internal/core/domains/dto"
	"io"
	"net/http"

	driverports "backend/internal/core/ports/driver_ports"
)

type CommentHandler struct {
	service driverports.CommentServiceInterface
	BaseHandler
}

func NewCommentHandler(service driverports.CommentServiceInterface, baseHandler BaseHandler) *CommentHandler {
	return &CommentHandler{service: service, BaseHandler: baseHandler}
}

func (commentHandler *CommentHandler) SubmitComment(w http.ResponseWriter, r *http.Request) {
	var size int64
	size = r.ContentLength
	if err := r.ParseMultipartForm(size); err != nil {
		commentHandler.handleError(w, r, 500, "fail", err)
		commentHandler.RenderError(w, 500, "fail")
		return
	}
	var comment dto.Comment
	var img []byte
	comment.PostID = r.Form["postID"][0]
	if len(r.FormValue("parentCommentID")) != 0 {
		comment.ParentID = r.Form["parentCommentID"][0]
	}
	comment.Content = r.Form["comment"][0]
	if len(r.FormValue("file")) != 0 {
		in, _, err := r.FormFile("file")
		if err != nil {
			commentHandler.handleError(w, r, 500, "fail", err)
			commentHandler.RenderError(w, 500, "fail")
			return
		}
		defer in.Close()
		img, err = io.ReadAll(in)
	}
	cookie, err := r.Cookie("session_id")
	comment.UserID = cookie.Value
	err = commentHandler.service.AddComment(comment, img, r.Context())
	if err != nil {
		commentHandler.handleError(w, r, 500, "fail", err)
		commentHandler.RenderError(w, 500, "fail")
		return
	}
}
