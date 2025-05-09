package drivenports

import (
	"backend/internal/core/domains/dao"
	"context"
)

type SessionRepoInterface interface {
	AddSession(ctx context.Context, session dao.Session) error
	GetSessionById(id string) (dao.Session, error)
}
