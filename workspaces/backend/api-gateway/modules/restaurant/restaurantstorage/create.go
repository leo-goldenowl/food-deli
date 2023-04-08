package restaurantstorage

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Create(
	ctx context.Context,
	data *restaurantmodel.RestaurantCreate,
) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
