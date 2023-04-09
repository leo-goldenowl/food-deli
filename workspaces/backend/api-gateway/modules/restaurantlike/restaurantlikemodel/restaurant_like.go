package restaurantlikemodel

import (
	"time"

	"github.com/google/uuid"
)

type RestaurantLike struct {
	RestaurantId uuid.UUID  `json:"restaurantId" gorm:"column:restaurant_id;type:uuid;not null"`
	UserId       uuid.UUID  `json:"userId" gorm:"column:user_id;type:uuid;not null"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"column:created_at;not null"`
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}
