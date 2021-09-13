package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	authclient "e.coding.net/double-j/ego/colago/samples/auth-client"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/user"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/convertor"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
	"golang.org/x/net/context"
)

func init() {
	ioc.AppendInjection(func(
		reRepo *repo.AccountUserRepo,
		userRepo *repo.UserRepo,
		accRepo *repo.AccountRepo,
		authcli *authclient.AuthClient,
	) user.UserGateway {
		return &UserGatewayImpl{
			reRepo:   reRepo,
			userRepo: userRepo,
			accRepo:  accRepo,
			authcli:  authcli,
		}
	})
}

type UserGatewayImpl struct {
	reRepo   *repo.AccountUserRepo
	userRepo *repo.UserRepo
	accRepo  *repo.AccountRepo
	authcli  *authclient.AuthClient
}

func (u *UserGatewayImpl) FindByAccount(ctx context.Context, dto *client.DTO, acc *account.Account) (*user.User, error) {
	rela, err := u.reRepo.FindByAccountId(dto.TenantId, acc.Id())
	if err != nil {
		return nil, err
	}
	usr, err := u.userRepo.FindById(dto.TenantId, rela.UserId)
	if err != nil {
		return nil, err
	}
	userEntity, err := convertor.PoToUserEntity(ctx, usr)
	if err != nil {
		return nil, err
	}
	userEntity.SetDto(dto)
	return userEntity, err
}

func (u *UserGatewayImpl) CreateUser(ctx context.Context, dto *client.DTO, user *user.User) error {
	userPo, err := convertor.EntityToUserPo(dto, user)
	if err != nil {
		return err
	}
	userPo, err = u.userRepo.InsertOne(userPo)
	if err != nil {
		return err
	}
	relations := make([]*po.RelationAccountUser, 0)
	for _, acc := range user.Accounts() {
		accountPo, err := convertor.EntityToAccountPo(dto, acc)
		if err != nil {
			return err
		}
		accountPo, err = u.accRepo.InsertOne(accountPo)
		if err != nil {
			return err
		}
		relations = append(relations, &po.RelationAccountUser{
			UserId:    userPo.ID,
			AccountId: accountPo.ID,
			TenantId:  dto.TenantId,
		})
	}
	_, err = u.reRepo.InsertBatch(relations)
	if err != nil {
		return err
	}
	return u.authcli.CreateRoleAuthCodes(ctx, &client.CreateAuthCmd{
		Dto:    dto,
		UserId: userPo.ID,
		Roles:  user.Roles(),
		Auths:  user.Auths(),
	})
}
