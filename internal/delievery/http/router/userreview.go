package v1

import (
	"github.com/gin-gonic/gin"
	a "legocy-go/internal/app"
	"legocy-go/internal/delievery/http/handlers/marketplace"
	"legocy-go/internal/delievery/http/middleware"
)

func (r V1router) addUserReviews(
	rg *gin.RouterGroup,
	app *a.App) {

	handler := marketplace.NewUserReviewHandler(app.GetUserReviewService())

	items := rg.Group("/user-reviews").Use(v1.Auth())
	{
		items.GET("/", handler.ListUserReviews)
		items.GET("/:reviewID", handler.UserReviewDetail)

		{
			items.POST("/", handler.CreateUserReview)
		}
		items.Use(v1.ReviewOwnerOrAdmin("reviewId", app.GetUserReviewRepo()))
		{
			items.DELETE("/:reviewId", handler.DeleteUserReview)
		}
	}
}
