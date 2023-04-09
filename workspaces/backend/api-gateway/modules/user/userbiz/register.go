package userbiz

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/user/usermodel"
)

type RegisterStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
	CreateUser(
		ctx context.Context,
		data *usermodel.UserCreate,
	) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	registerStore RegisterStore
	hasher        Hasher
}

func NewRegisterBiz(registerStore RegisterStore, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStore: registerStore,
		hasher:        hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) (*usermodel.User, error) {
	user, _ := biz.registerStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return nil, usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code
	data.Status = 1

	if err := biz.registerStore.CreateUser(ctx, data); err != nil {
		return nil, common.ErrDB(err)
	}

	newUser, err := biz.registerStore.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, common.ErrDB(err)
	}
	
	return newUser, nil
}
