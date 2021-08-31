package role

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
)

type RoleGateway interface {
	FindByUserId(dto *client.DTO, userId uint64) ([]*Role, error)
}
