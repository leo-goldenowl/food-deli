package main

import (
	"api-gateway/src/component"
	"api-gateway/src/database"
	"api-gateway/src/modules/restaurant/restaurantmodel"
	"api-gateway/src/modules/restaurant/restauranttransport/ginrestaurant"

	"log"
	"net/http"

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
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)

	// CRUD

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

		restaurants.PUT("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			var data restaurantmodel.RestaurantUpdate

			if err := ctx.ShouldBindJSON(&data); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid request body",
				})
				return
			}

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "cannot update restaurant",
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "restaurant updated",
			})
		})

		restaurants.DELETE("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			if err := db.Where("id = ?", id).Delete(&restaurantmodel.Restaurant{}).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "cannot delete restaurant",
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "restaurant deleted",
			})
		})
	}

	return r.Run()
}
