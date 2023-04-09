package userstorage

import (
	"context"
	"errors"

	"api-gateway/common"
	"api-gateway/modules/user/usermodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result *usermodel.User

	if err := db.Where(conditions).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return result, nil
}
