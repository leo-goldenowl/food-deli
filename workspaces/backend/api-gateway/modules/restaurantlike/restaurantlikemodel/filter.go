package restaurantlikemodel

import "github.com/google/uuid"

type Filter struct {
	RestaurantId uuid.UUID  `json:"-" form:"restaurant_id"`
	UserId       uuid.UUID `json:"-" form:"user_id"`
}
