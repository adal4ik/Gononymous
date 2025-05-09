package driverports

import (
	"context"

	"backend/internal/core/domains/dao"
)

type SessionServiceDriverInterface interface {
	CreateSession(ctx context.Context) (string, error)
	GetSessionById(id string,ctx context.Context) (dao.Session, error)
}
