package ginupload

import (
	"net/http"

	"api-gateway/common"
	"api-gateway/component"
	"api-gateway/modules/upload/uploadbiz"
	"api-gateway/modules/upload/uploadstorage"

	"github.com/gin-gonic/gin"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Save file to local
		// ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))

		folder := ctx.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), imgStore)
		img, err := biz.Upload(ctx.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
