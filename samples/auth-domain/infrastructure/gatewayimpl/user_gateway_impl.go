package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/user"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"golang.org/x/net/context"
)

func init() {
	ioc.AppendInjection(func(roleRepo *repo.UserRoleRepo, authRepo *repo.UserAuthRepo) user.UserGateway {
		return &UserGatewayImpl{
			roleRepo: roleRepo,
			authRepo: authRepo,
		}
	})
}

type UserGatewayImpl struct {
	roleRepo *repo.UserRoleRepo
	authRepo *repo.UserAuthRepo
}

func (u *UserGatewayImpl) SaveRoleAuth(ctx context.Context, dto *client.DTO, user *user.User) error {
	rolePos, err := convertor.EntitysToBatchRolePos(dto, user.Id(), user.Roles())
	if err != nil {
		return err
	}
	_, err = u.roleRepo.InsertBatch(rolePos)
	if err != nil {
		return err
	}
	authPos, err := convertor.EntitysToBatchAuthPos(dto, user.Id(), user.Auths())
	if err != nil {
		return err
	}
	_, err = u.authRepo.InsertBatch(authPos)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserGatewayImpl) FindById(ctx context.Context, dto *client.DTO, userId uint64) (*user.User, error) {
	usr := new(user.User)
	usr.SetId(userId)
	usr.SetDto(dto)
	usr.SetCtx(ctx)
	return usr, nil
}
