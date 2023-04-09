package restaurantlikestorage

import (
	"context"
	"errors"

	"api-gateway/common"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindData(
	ctx context.Context,
	conditions map[string]interface{},
) (*restaurantlikemodel.RestaurantLike, error) {
	db := s.db

	var result *restaurantlikemodel.RestaurantLike

	if err := db.Where(conditions).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return result, nil
}
