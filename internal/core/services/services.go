package services

import (
	db "Gononymous/internal/adapters/driven/database"
	"Gononymous/internal/adapters/driven/s3"
	driverports "Gononymous/internal/core/ports/driver_ports"
)

type Service struct {
	PostsService driverports.PostDriverPortInterface
}

func New(repo *db.Repository) *Service {
	return &Service{
		PostsService: NewPostService(repo.PostRepo, s3.NewS3ImageCollector()),
	}
}
