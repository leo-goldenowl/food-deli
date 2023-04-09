package restaurantlikebiz

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"

	"github.com/google/uuid"
)

type UserLikeRestaurantStore interface {
	FindData(
		ctx context.Context,
		conditions map[string]interface{},
	) (*restaurantlikemodel.RestaurantLike, error)
	Create(
		ctx context.Context,
		data *restaurantlikemodel.RestaurantLike,
	) error
}

type IncreaseLikedCountStore interface {
	IncreaseLikedCount(ctx context.Context, id uuid.UUID) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncreaseLikedCountStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore IncreaseLikedCountStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, incStore: incStore}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.RestaurantLike,
) error {
	likedRestaurant, _ := biz.store.FindData(ctx, map[string]interface{}{"user_id": data.UserId, "restaurant_id": data.RestaurantId})

	if likedRestaurant != nil {
		return restaurantlikemodel.ErrLikedRestaurant()
	}

	err := biz.store.Create(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}
	go func() {
		defer common.AppRecover()
		_ = biz.incStore.IncreaseLikedCount(ctx, data.RestaurantId)
	}()

	return nil
}
