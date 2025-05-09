package services

import (
	"context"

	driverports "backend/internal/core/ports/driver_ports"
)

type UserService struct {
	UserRepo driverports.UserDriverPortInterface
}

func NewUserService(userRepo driverports.UserDriverPortInterface) *UserService {
	return &UserService{UserRepo: userRepo}
}
func (s *UserService) ChangeName(userId string, newName string, ctx context.Context) error {
	return s.UserRepo.ChangeName(userId, newName, ctx)
}
