package driverports

import "context"

type SessionServiceDriverInterface interface {
	CreateSession(ctx context.Context) (string, error)
}
