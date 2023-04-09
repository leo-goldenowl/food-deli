package restaurantlikemodel

import (
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
