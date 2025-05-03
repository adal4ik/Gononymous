package handlers

import driverports "backend/internal/core/ports/driver_ports"

type CommentHandler struct {
	service driverports.PostDriverPortInterface
	BaseHandler
}

func NewCommentHandler(service driverports.PostDriverPortInterface, baseHandler BaseHandler) *CommentHandler {
	return &CommentHandler{service: service, BaseHandler: baseHandler}
}
