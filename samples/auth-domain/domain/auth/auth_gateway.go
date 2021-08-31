package auth

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
)

type AuthGateway interface {
	FindByUserId(dto *client.DTO, userId uint64) ([]*Auth, error)
}
