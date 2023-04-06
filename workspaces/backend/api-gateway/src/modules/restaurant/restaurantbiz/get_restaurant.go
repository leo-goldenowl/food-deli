package restaurantbiz

import (
	"context"

	"api-gateway/src/modules/restaurant/restaurantmodel"

	"github.com/google/uuid"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(
	ctx context.Context,
	id uuid.UUID,
) (restaurantmodel.Restaurant, error) {	
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	return data, err
}
