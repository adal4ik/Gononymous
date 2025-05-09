package handlers

import (
	"html/template"
	"log/slog"
	"net/http"

	"backend/internal/core/services"
	"backend/utils"
)

type BaseHandler struct {
	logger slog.Logger
}

func NewBaseHandler(logger slog.Logger) *BaseHandler {
	return &BaseHandler{logger: logger}
}

type Handler struct {
	PostHandler    *PostsHandler
	CatalogHandler *CatalogHandler
	CommentHandler *CommentHandler
	UserHandler    *UserHandler
	ArchiveHandler *ArchiveHandler
}

func (b *BaseHandler) handleError(w http.ResponseWriter, r *http.Request, code int, message string, err error) {
	if err != nil {
		b.logger.Error(message, "error", err, "code", code, "url", r.URL.Path)
	} else {
		b.logger.Error(message, "code", code, "url", r.URL.Path)
	}

	jsonErr := utils.APIError{
		Code:     code,
		Message:  message,
		Resource: r.URL.Path,
	}
	jsonErr.Send(w)
}

func (b *BaseHandler) RenderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("web/templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	data := struct {
		Code    int
		Message string
	}{
		Code:    code,
		Message: message,
	}

	tmpl.Execute(w, data)
}

func New(service *services.Service, baseHandler BaseHandler) *Handler {
	return &Handler{
		PostHandler:    NewPostHandler(service.PostsService, service.CommentService, service.SessionService, baseHandler),
		CatalogHandler: NewCatalogHandler(service.PostsService, baseHandler),
		CommentHandler: NewCommentHandler(service.CommentService, baseHandler),
		UserHandler:    NewUserHandler(service.UserService, baseHandler),
		ArchiveHandler: NewArchiveHandler(service.PostsService, baseHandler, service.CommentService, service.SessionService),
	}
}
