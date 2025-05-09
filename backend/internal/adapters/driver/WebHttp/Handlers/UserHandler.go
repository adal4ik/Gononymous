package handlers

import driverports "backend/internal/core/ports/driver_ports"

type UserHandler struct {
	service driverports.UserDriverPortInterface
	BaseHandler
}

func NewUserHandler(service driverports.UserDriverPortInterface, baseHandler BaseHandler) *UserHandler {
	return &UserHandler{
		service:     service,
		BaseHandler: baseHandler,
	}
}
