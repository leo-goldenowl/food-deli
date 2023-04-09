package restaurantlikestorage

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) Create(
	ctx context.Context,
	data *restaurantlikemodel.RestaurantLike,
) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
