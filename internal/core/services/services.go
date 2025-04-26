package services

import (
	db "Gononymous/internal/adapters/driven/database"
	driverports "Gononymous/internal/core/ports/driver_ports"
)

type Service struct {
	PostsService driverports.PostDriverPortInterface
}

func New(repo *db.Repository) *Service {
	return &Service{
		PostsService: NewPostService(repo.PostRepo),
	}
}
