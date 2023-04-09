package uploadbiz

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"api-gateway/common"
	"api-gateway/component/uploadprovider"
	"api-gateway/modules/upload/uploadmodel"
)

type CreateImageStore interface {
	// CreateImage(ctx context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStore
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStore) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	// fileBytes := bytes.NewReader(data)

	// w, h, err := getImageDimension(fileBytes)
	// if err != nil {
	// 	return nil, uploadmodel.ErrFileIsNotImage(err)
	// }

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUpload(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	// img.Width = w
	// img.Height = h
	img.CloudName = "s3"
	img.Extension = fileExt

	return img, err
}

// func getImageDimension(reader io.Reader) (int, int, error) {
// 	img, _, err := image.Decode(reader)
// 	if err != nil {
// 		log.Println("err: ", err)
// 		return 0, 0, err
// 	}

// 	fmt.Println(img.Bounds().Dx(), img.Bounds().Dy())

// 	return 0, 0, nil
// }
