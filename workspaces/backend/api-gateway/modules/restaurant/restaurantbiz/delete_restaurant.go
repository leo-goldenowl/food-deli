package restaurantbiz

import (
	"context"
	"fmt"

	"api-gateway/common"
	"api-gateway/modules/restaurant/restaurantmodel"

	"github.com/google/uuid"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	SoftDeleteData(ctx context.Context, id uuid.UUID) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id uuid.UUID) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	fmt.Println("oldData", oldData)

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}

	return nil
}
