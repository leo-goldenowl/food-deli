package restaurantstorage

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurant/restaurantmodel"

	"github.com/google/uuid"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id uuid.UUID,
) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
