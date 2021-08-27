package convertor

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/model/entity"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
)

func ToAccountEntity(account *po.Account) (*entity.Account, error) {
	accountBean, err := domain.GetDomainFactory().Create("entity.Account")
	if err != nil {
		return nil, err
	}
	accountEntity := accountBean.(*entity.Account)
	accountEntity.SetId(account.ID)
	accountEntity.SetAccType(account.AccType)
	accountEntity.SetAccKey(account.AccKey)
	accountEntity.SetPassword(account.Password)
	accountEntity.SetEnable(account.Enable == 1)
	return accountEntity, nil
}
