package account

import (
	"e.coding.net/double-j/ego/colago/samples/shared/client"
)

type AccountGateway interface {
	FindAccountByAccKey(dto *client.DTO, accKey string) (*Account, error)
}
