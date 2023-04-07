package restaurantbiz

import (
	"context"
	"errors"

	"api-gateway/modules/restaurant/restaurantmodel"

	"github.com/google/uuid"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateData(
		ctx context.Context,
		id uuid.UUID,
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(
	ctx context.Context,
	id uuid.UUID,
	data *restaurantmodel.RestaurantUpdate,
) (*restaurantmodel.Restaurant, error) {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	if oldData.Status == 0 {
		return nil, errors.New("restaurant deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return nil, err
	}

	newData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return newData, nil
}
