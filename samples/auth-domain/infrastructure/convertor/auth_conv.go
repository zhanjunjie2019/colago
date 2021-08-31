package convertor

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/domain/auth"
	"e.coding.net/double-j/ego/colago/samples/auth-domain/infrastructure/repo/po"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
)

func CodesToBatchAuthEntitys(codes []string) ([]*auth.Auth, error) {
	authEntitys := make([]*auth.Auth, 0)
	for _, c := range codes {
		authBean, err := domain.GetDomainFactory().Create("auth.Auth")
		if err != nil {
			return nil, err
		}
		authEntity := authBean.(*auth.Auth)
		authEntity.SetCode(c)
		authEntitys = append(authEntitys, authEntity)
	}
	return authEntitys, nil
}

func PosToBatchAuthEntitys(auths []*po.RelationUserAuth) ([]*auth.Auth, error) {
	authEntitys := make([]*auth.Auth, 0)
	for _, a := range auths {
		entity, err := PoToAuthEntity(a)
		if err != nil {
			return nil, err
		}
		authEntitys = append(authEntitys, entity)
	}
	return authEntitys, nil
}

func PoToAuthEntity(a *po.RelationUserAuth) (*auth.Auth, error) {
	authBean, err := domain.GetDomainFactory().Create("auth.Auth")
	if err != nil {
		return nil, err
	}
	authEntity := authBean.(*auth.Auth)
	authEntity.SetCode(a.AuthCode)
	return authEntity, nil
}

func EntitysToBatchAuthPos(dto *client.DTO, userid uint64, auths []*auth.Auth) ([]*po.RelationUserAuth, error) {
	pos := make([]*po.RelationUserAuth, 0)
	for _, a := range auths {
		authPo, err := EntityToAuthPo(dto, userid, a)
		if err != nil {
			return nil, err
		}
		pos = append(pos, authPo)
	}
	return pos, nil
}

func EntityToAuthPo(dto *client.DTO, userid uint64, auth *auth.Auth) (*po.RelationUserAuth, error) {
	return &po.RelationUserAuth{
		UserId:   userid,
		AuthCode: auth.Code(),
		TenantId: dto.TenantId,
	}, nil
}
