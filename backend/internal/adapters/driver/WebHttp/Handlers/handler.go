package handlers

import (
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

func New(service *services.Service, baseHandler BaseHandler) *Handler {
	return &Handler{
		PostHandler:    NewPostHandler(service.PostsService, baseHandler),
		CatalogHandler: NewCatalogHandler(service.PostsService),
	}
}
