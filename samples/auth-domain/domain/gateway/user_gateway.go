package gateway

import (
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/model/entity"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
)

type UserGateway interface {
	FindById(dto *client.DTO, userId uint64) (*entity.User, error)
}
