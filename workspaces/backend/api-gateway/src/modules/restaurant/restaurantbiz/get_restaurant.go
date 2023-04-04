package restaurantbiz

import (
	"api-gateway/src/modules/restaurant/restaurantmodel"
	"context"

	"github.com/google/uuid"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
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
) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	return result, err
}
