package handlers

import (
	"log/slog"

	"Gononymous/internal/core/services"
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

func New(service *services.Service) *Handler {
	return &Handler{
		PostHandler:    NewPostHandler(service.PostsService),
		CatalogHandler: NewCatalogHandler(),
	}
}
