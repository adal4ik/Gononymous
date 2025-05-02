package services

import (
	db "backend/internal/adapters/driven/database"
	"backend/internal/adapters/driven/s3"
	driverports "backend/internal/core/ports/driver_ports"
)

type Service struct {
	PostsService driverports.PostDriverPortInterface
}

func New(repo *db.Repository) *Service {
	return &Service{
		PostsService: NewPostService(repo.PostRepo, s3.NewS3ImageCollector()),
	}
}
