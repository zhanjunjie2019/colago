package convertor

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/samples/shared/client"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/account"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/user"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
	"time"
)

func UserCreateDtoToUserEntity(cmd *client.CreateUserCmd) (*user.User, error) {
	userBean, err := domain.GetDomainFactory().Create("user.User")
	if err != nil {
		return nil, err
	}
	accountEntity, err := UserCreateDtoToAccountEntity(cmd)
	if err != nil {
		return nil, err
	}
	userEntity := userBean.(*user.User)
	userEntity.SetAccounts([]*account.Account{accountEntity})
	userEntity.SetFirstName(cmd.FirstName)
	userEntity.SetLastName(cmd.LastName)
	userEntity.SetAge(uint8(cmd.Age))
	userEntity.SetBirthday(time.Unix(int64(cmd.BirthdayTs), 0))
	userEntity.SetEmail(cmd.Email)
	userEntity.SetPhoneNumber(cmd.PhoneNumber)
	userEntity.SetStatus(1)
	userEntity.SetRoles(cmd.Roles)
	userEntity.SetAuths(cmd.Auths)
	userEntity.SetDto(cmd.Dto)
	return userEntity, nil
}

func PoToUserEntity(u *po.UserInfo) (*user.User, error) {
	userBean, err := domain.GetDomainFactory().Create("user.User")
	if err != nil {
		return nil, err
	}
	userEntity := userBean.(*user.User)
	userEntity.SetId(u.ID)
	userEntity.SetFirstName(u.FirstName)
	userEntity.SetLastName(u.LastName)
	userEntity.SetAge(u.Age)
	userEntity.SetBirthday(u.Birthday)
	userEntity.SetEmail(u.Email)
	userEntity.SetPhoneNumber(u.PhoneNumber)
	userEntity.SetStatus(u.Status)
	return userEntity, nil
}

func EntityToUserPo(dto *client.DTO, user *user.User) (*po.UserInfo, error) {
	userPo := new(po.UserInfo)
	userPo.TenantId = dto.TenantId
	userPo.FirstName = user.FirstName()
	userPo.LastName = user.LastName()
	userPo.Age = user.Age()
	userPo.Birthday = user.Birthday()
	userPo.Email = user.Email()
	userPo.PhoneNumber = user.PhoneNumber()
	userPo.Status = user.Status()
	return userPo, nil
}
