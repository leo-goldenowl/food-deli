package restaurantstorage

import (
	"context"

	"api-gateway/modules/restaurant/restaurantmodel"

	"github.com/google/uuid"
)

func (s *sqlStore) UpdateData(ctx context.Context, id uuid.UUID, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
