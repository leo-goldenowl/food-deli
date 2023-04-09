package restaurantlikebiz

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"

	"github.com/google/uuid"
)

type UserUnlikeRestaurantStore interface {
	FindData(
		ctx context.Context,
		conditions map[string]interface{},
	) (*restaurantlikemodel.RestaurantLike, error)
	Delete(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type DecreaseLikedCountStore interface {
	DecreaseLikedCount(ctx context.Context, id uuid.UUID) error
}

type UserUnlikeRestaurantBiz struct {
	store    UserUnlikeRestaurantStore
	decStore DecreaseLikedCountStore
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore, decStore DecreaseLikedCountStore) *UserUnlikeRestaurantBiz {
	return &UserUnlikeRestaurantBiz{store: store, decStore: decStore}
}

func (biz *UserUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.RestaurantLike,
) error {
	conditions := map[string]interface{}{"user_id": data.UserId, "restaurant_id": data.RestaurantId}

	likedRestaurant, _ := biz.store.FindData(ctx, conditions)

	if likedRestaurant == nil {
		return restaurantlikemodel.ErrHasNotLikedRestaurant()
	}

	err := biz.store.Delete(ctx, conditions)
	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		_ = biz.decStore.DecreaseLikedCount(ctx, data.RestaurantId)
	}()

	return nil
}
