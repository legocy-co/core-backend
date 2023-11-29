package v1

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	"legocy-go/internal/app"
	"legocy-go/internal/delivery/http/handlers/users"
	"legocy-go/internal/delivery/http/handlers/users/auth"
	"legocy-go/internal/delivery/http/handlers/users/userImage"
	"legocy-go/internal/delivery/http/middleware"
	jwt "legocy-go/pkg/auth/jwt/middleware"
)

func (r V1router) addUsers(rg *gin.RouterGroup, app *app.App) {

	// Authentication

	authHandler := auth.NewTokenHandler(app.GetUserService())

	authRouter := rg.Group("/users/auth")
	{
		authRouter.POST("/sign-in", authHandler.GenerateToken)
		authRouter.POST("/refresh", authHandler.RefreshToken)
		authRouter.POST("/register", authHandler.UserRegister)
	}

	// User Profile

	profileHandler := users.NewUserProfilePageHandler(
		app.GetMarketItemService(), app.GetUserService(), app.GetUserReviewService(), app.GetUserImagesService())

	profileRoutes := rg.Group("/users/profile").Use(jwt.IsAuthenticated())
	{
		profileRoutes.GET("/:userID", profileHandler.UserProfilePageDetail)
	}

	// User Reviews

	reviewsHandler := users.NewUserReviewHandler(app.GetUserReviewService())
	reviewsRoutes := rg.Group("/users/reviews").Use(jwt.IsAuthenticated())
	{
		reviewsRoutes.GET("/", reviewsHandler.ListUserReviews)
		reviewsRoutes.GET("/:reviewID", reviewsHandler.UserReviewDetail)

		{
			reviewsRoutes.POST("/", reviewsHandler.CreateUserReview)
		}
		reviewsRoutes.Use(middleware.ReviewOwnerOrAdmin("reviewId", app.GetUserReviewRepo()))
		{
			reviewsRoutes.DELETE("/:reviewId", reviewsHandler.DeleteUserReview)
		}
	}

	// User Images

	userImagesHandler := userImage.NewUserImageHandler(
		app.GetUserImagesService(),
		app.GetImageStorageClient())

	userImages := rg.Group("/users/images")
	{
		userImages.GET("/download", userImagesHandler.DownloadImage)
		userImages.Use(jwt.IsOwnerOrAdmin("userID"))
		{
			userImages.GET("/:userID", userImagesHandler.ListImages)
			userImages.POST("/:userID", userImagesHandler.UploadUserImage)
		}
	}
}
