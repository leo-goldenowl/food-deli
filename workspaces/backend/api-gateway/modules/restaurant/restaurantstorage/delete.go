package restaurantstorage

import (
	"api-gateway/modules/restaurant/restaurantmodel"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id uuid.UUID,
) error {
	db := s.db

	fmt.Println("hihi 1", id)

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}

	fmt.Println("hihi 2")

	return nil
}
