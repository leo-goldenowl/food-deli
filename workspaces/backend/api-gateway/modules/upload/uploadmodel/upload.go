package uploadmodel

import (
	"api-gateway/common"
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrUploadFile")
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save file",
		"ErrUploadFile")
}
