package site

import (
	"github.com/gin-gonic/gin"
	_ "github.com/legocy-co/legocy/docs"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/users"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/users/auth"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/users/auth/google"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/users/profile"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/users/userImage"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware"
	"github.com/legocy-co/legocy/internal/pkg/app"
	jwt "github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddUsers(rg *gin.RouterGroup, app *app.App) {

	// Authentication
	authHandler := auth.NewTokenHandler(app.GetUserService())
	authGoogleHandler := google.NewHandler(app)

	authRouter := rg.Group("/users/auth")
	{
		authRouter.POST("/sign-in", authHandler.GenerateToken)
		authRouter.POST("/refresh", authHandler.RefreshToken)
		authRouter.POST("/register", authHandler.UserRegister)

		// Google
		authRouter.POST("/google/sign-in", authGoogleHandler.SignIn)
		authRouter.POST("/google/sign-up", authGoogleHandler.SignUp)
	}

	// User Profile

	profileHandler := profile.NewUserProfilePageHandler(
		app.GetMarketItemService(), app.GetUserService(), app.GetUserReviewService())

	profileRoutes := rg.Group("/users/profile").Use(jwt.IsAuthenticated())
	{

		profileRoutes.GET("/", profileHandler.CurrentUserProfilePage)
		profileRoutes.GET("/header/", profileHandler.CurrentUserProfileHeader)
		profileRoutes.GET("/:userID", profileHandler.UserProfilePageDetail)

		privateProfileRoutes := profileRoutes.Use(middleware.IsOwnerOrAdmin("userID"))
		{
			privateProfileRoutes.PUT("/:userID", profileHandler.UpdateUserProfile)
		}

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
		app.GetImageStorageClient(),
	)

	userImages := rg.Group("/users/images").Use(jwt.IsOwnerOrAdmin("userID"))
	{
		userImages.GET("/:userID", userImagesHandler.ListImages)
		userImages.POST("/:userID/avatar", userImagesHandler.UploadUserImage)
	}
}
