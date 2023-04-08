package ginrestaurant

import (
	"net/http"

	"api-gateway/common"
	"api-gateway/component"
	"api-gateway/modules/restaurant/restaurantbiz"
	"api-gateway/modules/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		result, err := biz.GetRestaurant(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
