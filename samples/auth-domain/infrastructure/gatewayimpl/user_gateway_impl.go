package gatewayimpl

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/model/entity"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
)

func init() {
	_ = ioc.InjectSimpleBean(new(UserGatewayImpl))
}

type UserGatewayImpl struct {
}

func (u *UserGatewayImpl) New() ioc.AbsBean {
	return u
}

func (u *UserGatewayImpl) FindById(dto *client.DTO, userId uint64) (*entity.User, error) {
	userBean, err := domain.GetDomainFactory().Create("entity.User")
	if err != nil {
		return nil, err
	}
	user := userBean.(*entity.User)
	user.SetId(userId)
	return user, nil
}
