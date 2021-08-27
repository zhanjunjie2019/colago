package gateway

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/model/entity"
)

type AccountGateway interface {
	FindAccountByAccKey(dto *client.DTO, accKey string) (*entity.Account, error)
}
