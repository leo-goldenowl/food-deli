package restaurantstorage

import (
	"context"
	"errors"

	"api-gateway/modules/restaurant/restaurantmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	db := s.db

	var result *restaurantmodel.Restaurant

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	err := db.Where(conditions).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("data not found")
	}

	return result, nil
}
