package drivenports

import (
	"context"

	"backend/internal/core/domains/dao"
)

type SessionRepoInterface interface {
	AddSession(ctx context.Context, session dao.Session) error
	GetSessionById(id string,ctx context.Context) (dao.Session, error)
}
