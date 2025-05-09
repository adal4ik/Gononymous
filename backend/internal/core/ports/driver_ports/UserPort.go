package driverports

import "context"

type UserDriverPortInterface interface {
	ChangeName(userId string, newName string, ctx context.Context) error
}
