package user

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
)

type UserGateway interface {
	SaveRoleAuth(dto *client.DTO, user *User) error
	FindById(dto *client.DTO, userId uint64) (*User, error)
}
