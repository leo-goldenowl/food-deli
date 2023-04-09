package restaurantlikestorage

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"

	"github.com/google/uuid"
)

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]int, error) {
	result := make(map[uuid.UUID]int)

	type sqlData struct {
		RestaurantId uuid.UUID `gorm:"column:restaurant_id"`
		LikedCount   int       `gorm:"column:count"`
	}

	var listLike []sqlData

	if err := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikedCount
	}

	return result, nil
}
