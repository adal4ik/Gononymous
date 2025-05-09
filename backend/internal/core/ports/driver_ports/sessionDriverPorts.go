package driverports

import (
	"backend/internal/core/domains/dao"
	"context"
)

type SessionServiceDriverInterface interface {
	CreateSession(ctx context.Context) (string, error)
	GetSessionById(id string) (dao.Session, error)
}
