package ginuser

import (
	"net/http"
	"strings"

	"api-gateway/common"
	"api-gateway/component"
	"api-gateway/component/hasher"
	"api-gateway/component/tokenprovider/jwt"
	"api-gateway/modules/user/userbiz"
	"api-gateway/modules/user/usermodel"
	"api-gateway/modules/user/userstorage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := ctx.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if len(strings.TrimSpace(loginUserData.Email)) == 0 || len(strings.TrimSpace(loginUserData.Password)) == 0 {
			panic(usermodel.ErrUsernameOrPasswordInvalid)
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(ctx.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
