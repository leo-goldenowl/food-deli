package restaurantbiz

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurant/restaurantmodel"

	"github.com/google/uuid"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type RestaurantLikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]int, error)
}

type listRestaurantBiz struct {
	store     ListRestaurantStore
	likeStore RestaurantLikeStore
}

func NewListRestaurantBiz(store ListRestaurantStore, likeStore RestaurantLikeStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	// ids := make([]uuid.UUID, len(result))

	// for i := range result {
	// 	ids[i] = result[i].Id
	// }

	// restaurantsLikes, _ := biz.likeStore.GetRestaurantLikes(ctx, ids)

	// if v := restaurantsLikes; v != nil {
	// 	for i, item := range result {
	// 		result[i].LikedCount = restaurantsLikes[item.Id]
	// 	}
	// }

	return result, nil
}
