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
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(UserGatewayImpl))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type UserGatewayImpl struct {
	ReRepo   *repo.AccountUserRepo  `ij:"repo.AccountUserRepo"`
	UserRepo *repo.UserRepo         `ij:"repo.UserRepo"`
	AccRepo  *repo.AccountRepo      `ij:"repo.AccountRepo"`
	Authcli  *authclient.AuthClient `ij:"authclient.AuthClient"`
}

func (u *UserGatewayImpl) New() ioc.AbsBean {
	return u
}

func (u *UserGatewayImpl) FindByAccount(dto *client.DTO, acc *account.Account) (*user.User, error) {
	rela, err := u.ReRepo.FindByAccountId(dto.TenantId, acc.Id())
	if err != nil {
		return nil, err
	}
	usr, err := u.UserRepo.FindById(dto.TenantId, rela.UserId)
	if err != nil {
		return nil, err
	}
	userEntity, err := convertor.PoToUserEntity(usr)
	if err != nil {
		return nil, err
	}
	userEntity.SetDto(dto)
	return userEntity, err
}

func (u *UserGatewayImpl) CreateUser(dto *client.DTO, user *user.User) error {
	userPo, err := convertor.EntityToUserPo(dto, user)
	if err != nil {
		return err
	}
	userPo, err = u.UserRepo.InsertOne(userPo)
	if err != nil {
		return err
	}
	relations := make([]*po.RelationAccountUser, 0)
	for _, acc := range user.Accounts() {
		accountPo, err := convertor.EntityToAccountPo(dto, acc)
		if err != nil {
			return err
		}
		accountPo, err = u.AccRepo.InsertOne(accountPo)
		if err != nil {
			return err
		}
		relations = append(relations, &po.RelationAccountUser{
			UserId:    userPo.ID,
			AccountId: accountPo.ID,
			TenantId:  dto.TenantId,
		})
	}
	_, err = u.ReRepo.InsertBatch(relations)
	if err != nil {
		return err
	}
	return u.Authcli.CreateRoleAuthCodes(&client.CreateAuthCmd{
		Dto:    dto,
		UserId: userPo.ID,
		Roles:  user.Roles(),
		Auths:  user.Auths(),
	})
}
