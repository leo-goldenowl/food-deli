package main

import (
	"log"
	"net/http"

	"api-gateway/component"
	"api-gateway/config"
	"api-gateway/database"
	"api-gateway/middleware"
	"api-gateway/modules/restaurant/restaurantmodel"
	"api-gateway/modules/restaurant/restauranttransport/ginrestaurant"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"
	"api-gateway/modules/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"api-gateway/modules/user/usermodel"
	"api-gateway/modules/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db := database.CreateInstance()

	db.AutoMigrate(&usermodel.User{})
	db.AutoMigrate(&restaurantmodel.Restaurant{})
	db.AutoMigrate(&restaurantlikemodel.RestaurantLike{})

	db = db.Debug()

	if err := runService(db); err != nil {
		log.Fatal("can not start the server.\n", err)
	}
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db, config.App.JwtSecret)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))
	}

	return r.Run()
}
