package gateway

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/model/entity"
)

type UserGateway interface {
	FindByAccount(dto *client.DTO, acc *entity.Account) (*entity.User, error)
}
