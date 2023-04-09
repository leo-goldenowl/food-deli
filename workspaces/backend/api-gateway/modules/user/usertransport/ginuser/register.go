package ginuser

import (
	"net/http"

	"api-gateway/common"
	"api-gateway/component"
	"api-gateway/component/hasher"
	"api-gateway/modules/user/userbiz"
	"api-gateway/modules/user/usermodel"
	"api-gateway/modules/user/userstorage"

	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data usermodel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(data))
	}
}
