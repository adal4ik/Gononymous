package drivenports

import (
	"Gononymous/internal/core/domains/dao"
	"context"
)

type SessionRepoInterface interface {
	AddSession(ctx context.Context, session dao.Session) error
}
