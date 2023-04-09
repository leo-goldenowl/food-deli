package ginrestaurantlike

import (
	"net/http"

	"api-gateway/common"
	"api-gateway/component"
	"api-gateway/modules/restaurant/restaurantstorage"
	"api-gateway/modules/restaurantlike/restaurantlikebiz"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"
	"api-gateway/modules/restaurantlike/restaurantlikestorage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.RestaurantLike{
			RestaurantId: id,
			UserId:       requester.GetUserId(),
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantlikestorage.NewSQLStore(db)
		incStore := restaurantstorage.NewSQLStore(db)
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		if err := biz.LikeRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(id))
	}
}
