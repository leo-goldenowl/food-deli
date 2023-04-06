package restaurantstorage

import (
	"context"

	"api-gateway/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions).First(&result)

	return result, nil
}
