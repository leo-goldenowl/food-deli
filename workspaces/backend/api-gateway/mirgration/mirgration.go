package main

import (
	"log"

	"api-gateway/database"
	"api-gateway/modules/restaurant/restaurantmodel"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"
	"api-gateway/modules/user/usermodel"
)

func main() {
	db := database.CreateInstance()

	db.AutoMigrate(
		&usermodel.User{},
		&restaurantmodel.Restaurant{},
		&restaurantlikemodel.RestaurantLike{},
	)

	log.Println("migrated data!")
}
