package restaurantlikemodel

import "github.com/google/uuid"

type Filter struct {
	RestaurantId uuid.UUID `json:"restaurantId" form:"restaurant_id"`
	UserId       uuid.UUID `json:"userId" form:"user_id"`
}
