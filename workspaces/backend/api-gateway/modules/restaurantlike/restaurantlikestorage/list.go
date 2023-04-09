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

func (s *sqlStore) GetUsersLikeRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	db := s.db

	var result []restaurantlikemodel.RestaurantLike

	db = db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId != uuid.Nil {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("User")

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User
	}

	return users, nil
}
