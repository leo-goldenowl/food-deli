package main

import (
	"fmt"
	"log"
	"net/http"

	"api-gateway/component"
	"api-gateway/component/uploadprovider"
	"api-gateway/config"
	"api-gateway/database"
	"api-gateway/middleware"
	"api-gateway/modules/restaurant/restauranttransport/ginrestaurant"
	"api-gateway/modules/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"api-gateway/modules/upload/uploadtransport/ginupload"
	"api-gateway/modules/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	s3Provider := uploadprovider.NewS3Provider(
		config.App.S3BucketName,
		config.App.S3Region,
		config.App.S3APIKey,
		config.App.S3SecretKey,
		config.App.S3Domain,
	)

	db := database.CreateInstance()
	
	if err := runService(db, s3Provider); err != nil {
		log.Fatal("can not start the server.\n", err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, config.App.JwtSecret, upProvider)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	// health
	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// core
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))

	// other routes
	v1.POST("/upload", middleware.RequireAuth(appCtx), ginupload.Upload(appCtx))

	// user
	users := v1.Group("/users", middleware.RequireAuth(appCtx))
	{
		users.GET("/me", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))
	}

	// restaurant
	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))
		restaurants.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appCtx))
		restaurants.DELETE("/:id/unlike", ginrestaurantlike.UserUnLikeRestaurant(appCtx))
	}

	return r.Run(fmt.Sprintf(":%s", config.App.Port))
}
