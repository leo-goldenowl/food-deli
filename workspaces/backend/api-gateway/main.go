package main

import (
	"log"
	"net/http"

	"api-gateway/component"
	"api-gateway/database"
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
		restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))

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
