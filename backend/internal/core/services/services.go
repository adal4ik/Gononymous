package services

import (
	db "backend/internal/adapters/driven/database"
	"backend/internal/adapters/driven/s3"
	driverports "backend/internal/core/ports/driver_ports"
)

type Service struct {
	PostsService   driverports.PostDriverPortInterface
	SessionService driverports.SessionServiceDriverInterface
	CommentService driverports.CommentServiceInterface
	UserService    driverports.UserDriverPortInterface
}

func New(repo *db.Repository) *Service {
	return &Service{
		PostsService:   NewPostService(repo.PostRepo, s3.NewS3ImageCollector()),
		SessionService: NewSessionService(repo.SessionRepo, repo.CharacterRepo),
		CommentService: NewCommentService(repo.CommentRepo, s3.NewS3ImageCollector()),
		UserService:    NewUserService(repo.UserRepo),
	}
}
