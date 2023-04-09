package component

import (
	"api-gateway/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, secretKey string, upProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, upProvider: upProvider}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}