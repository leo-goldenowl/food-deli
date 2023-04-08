package restaurantstorage

import (
	"context"
	"strings"

	"api-gateway/common"
	"api-gateway/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := s.db

	var result []restaurantmodel.Restaurant

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions).Where("status = ?", 1)

	if v := filter; v != nil {
		if len(strings.TrimSpace(v.Name)) > 0 {
			db = db.Where("name = ?", v.Name)
		}
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
