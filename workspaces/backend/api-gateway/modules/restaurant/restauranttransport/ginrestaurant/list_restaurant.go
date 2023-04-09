package ginrestaurant

import (
	"net/http"

	"api-gateway/common"
	"api-gateway/component"
	"api-gateway/modules/restaurant/restaurantbiz"
	"api-gateway/modules/restaurant/restaurantmodel"
	"api-gateway/modules/restaurant/restaurantstorage"
	"api-gateway/modules/restaurantlike/restaurantlikestorage"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter restaurantmodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		likeStore := restaurantlikestorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store, likeStore)

		result, err := biz.ListRestaurant(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
