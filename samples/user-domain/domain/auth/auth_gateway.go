package auth

import "e.coding.net/double-j/ego/colago/samples/shared/client"

type AuthGateway interface {
	FindRolesByUserId(dto *client.DTO, userId uint64) ([]string, error)
	FindAuthsByUserId(dto *client.DTO, userId uint64) ([]string, error)
}
