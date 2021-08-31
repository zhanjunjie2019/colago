package convertor

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
)

func UserCreateDtoToAccountEntity(cmd *client.CreateUserCmd) (*account.Account, error) {
	accountBean, err := domain.GetDomainFactory().Create("account.Account")
	if err != nil {
		return nil, err
	}
	accountEntity := accountBean.(*account.Account)
	accountEntity.SetAccType(account.AccountType(cmd.AccType))
	accountEntity.SetAccKey(cmd.AccKey)
	accountEntity.SetPassword(cmd.Password)
	accountEntity.SetEnable(true)
	return accountEntity, nil
}

func PoToAccountEntity(a *po.Account) (*account.Account, error) {
	accountBean, err := domain.GetDomainFactory().Create("account.Account")
	if err != nil {
		return nil, err
	}
	accountEntity := accountBean.(*account.Account)
	accountEntity.SetId(a.ID)
	accountEntity.SetAccType(a.AccType)
	accountEntity.SetAccKey(a.AccKey)
	accountEntity.SetPassword(a.Password)
	accountEntity.SetEnable(a.Enable == 1)
	return accountEntity, nil
}

func EntityToAccountPo(dto *client.DTO, acc *account.Account) (*po.Account, error) {
	accPo := new(po.Account)
	accPo.TenantId = dto.TenantId
	accPo.AccType = acc.AccType()
	accPo.AccKey = acc.AccKey()
	accPo.Password = acc.Password()
	if acc.Enable() {
		accPo.Enable = 1
	} else {
		accPo.Enable = 0
	}
	return accPo, nil
}
