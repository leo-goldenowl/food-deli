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

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(id))
	}
}
