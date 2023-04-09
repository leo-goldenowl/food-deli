package ginrestaurantlike

import (
	"net/http"

	"api-gateway/common"
	"api-gateway/component"
	"api-gateway/modules/restaurantlike/restaurantlikebiz"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"
	"api-gateway/modules/restaurantlike/restaurantlikestorage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: id,
		}

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetMainDBConnection()
		store := restaurantlikestorage.NewSQLStore(db)
		biz := restaurantlikebiz.NewListUsersLikeRestaurantBiz(store)

		result, err := biz.ListUsers(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
