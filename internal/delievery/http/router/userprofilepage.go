package v1

import (
	"github.com/gin-gonic/gin"
	a "legocy-go/internal/app"
	"legocy-go/internal/delievery/http/handlers/marketplace"
	v1 "legocy-go/internal/delievery/http/middleware"
)

func (r V1router) addUserProfilePages(
	rg *gin.RouterGroup,
	app *a.App) {

	handler := marketplace.NewUserProfilePageHandler(
		app.GetMarketItemService(), app.GetUserService(), app.GetUserReviewService(),
		app.GetUserImagesService(), app.GetNotifyEventClient())

	items := rg.Group("/user-profile-pages/").Use(v1.Auth())
	{
		items.GET("/:userID", handler.UserProfilePageDetail)
	}
}
