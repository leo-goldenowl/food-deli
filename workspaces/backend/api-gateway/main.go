package main

import (
	"log"
	"net/http"

	"api-gateway/component"
	"api-gateway/database"
	"api-gateway/middleware"
	"api-gateway/modules/restaurant/restaurantmodel"
	"api-gateway/modules/restaurant/restauranttransport/ginrestaurant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db := database.CreateInstance()

	db.AutoMigrate(&restaurantmodel.Restaurant{})

	if err := runService(db); err != nil {
		log.Fatal("can not start the server.\n", err)
	}
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run()
}
