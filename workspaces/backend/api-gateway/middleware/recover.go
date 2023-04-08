package middleware

import (
	"api-gateway/common"
	"api-gateway/component"

	"github.com/gin-gonic/gin"
)

func Recover(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				appErr := common.ErrInternal(err.(error))
				ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		ctx.Next()
	}
}
