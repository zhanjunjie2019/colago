package convertor

import (
	"e.coding.net/double-j/ego/colago/common/domain"
	"e.coding.net/double-j/ego/colago/samples/user-domain/domain/model/entity"
	"e.coding.net/double-j/ego/colago/samples/user-domain/infrastructure/repo/po"
)

func ToUserEntity(user *po.UserInfo) (*entity.User, error) {
	userBean, err := domain.GetDomainFactory().Create("entity.User")
	if err != nil {
		return nil, err
	}
	userEntity := userBean.(*entity.User)
	userEntity.SetId(user.ID)
	userEntity.SetFirstName(user.FirstName)
	userEntity.SetLastName(user.LastName)
	userEntity.SetAge(user.Age)
	userEntity.SetBirthday(user.Birthday)
	userEntity.SetEmail(user.Email)
	userEntity.SetPhoneNumber(user.PhoneNumber)
	userEntity.SetStatus(user.Status)
	return userEntity, nil
}
