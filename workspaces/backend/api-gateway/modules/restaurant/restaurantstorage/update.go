package restaurantstorage

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurant/restaurantmodel"

	"github.com/google/uuid"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id uuid.UUID,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
