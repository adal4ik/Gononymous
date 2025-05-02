package drivenports

import (
	"context"

	"backend/internal/core/domains/dao"
)

type SessionRepoInterface interface {
	AddSession(ctx context.Context, session dao.Session) error
}
