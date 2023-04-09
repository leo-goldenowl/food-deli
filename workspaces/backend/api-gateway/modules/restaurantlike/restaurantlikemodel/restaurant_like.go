package restaurantlikemodel

import (
	"fmt"
	"time"

	"api-gateway/common"

	"github.com/google/uuid"
)

const EntityName = "RestaurantLike"

type RestaurantLike struct {
	RestaurantId uuid.UUID          `json:"restaurantId" gorm:"column:restaurant_id;type:uuid;not null"`
	UserId       uuid.UUID          `json:"userId" gorm:"column:user_id;type:uuid;not null"`
	CreatedAt    *time.Time         `json:"createdAt" gorm:"column:created_at;not null"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(err,
		"cannot like this restaurant",
		fmt.Sprintf("ErrCannotLike%s", EntityName))
}

func ErrLikedRestaurant() *common.AppError {
	return common.NewCustomError(nil,
		"liked this restaurant",
		fmt.Sprintf("ErrCannotLike%s", EntityName))
}

func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(err,
		"cannot unlike this restaurant",
		fmt.Sprintf("ErrCannotUnlike%s", EntityName))
}

func ErrHasNotLikedRestaurant() *common.AppError {
	return common.NewCustomError(nil,
		"has not liked this restaurant",
		fmt.Sprintf("ErrCannotUnlike%s", EntityName))
}
