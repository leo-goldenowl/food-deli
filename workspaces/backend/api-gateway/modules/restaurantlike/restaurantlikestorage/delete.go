package restaurantlikestorage

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) Delete(
	ctx context.Context,
	conditions map[string]interface{},
) error {
	db := s.db

	if err := db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Where(conditions).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
