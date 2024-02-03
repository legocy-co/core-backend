package site

import (
	"github.com/gin-gonic/gin"
	a "github.com/legocy-co/legocy/internal/app"
	collection "github.com/legocy-co/legocy/internal/delivery/http/handlers/collection/user_collection"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware"
)

func AddUserCollections(rg *gin.RouterGroup, app *a.App) {
	handler := collection.NewUserLegoCollectionHandler(app.GetUserCollectionService())

	userCollection := rg.Group("/collections")
	{
		userCollection.GET("/", handler.GetUserCollection)
		userCollection.POST("/", handler.AddLegoSetToUserCollection)
		userCollection.GET("/calculator/", handler.GetUserCollectionValuation)
		userCollection.Use(
			middleware.CollectionSetOwnerOrAdmin("setID", app.GetUserLegoSetsRepository()))
		{
			userCollection.PUT("/:setID", handler.UpdateUserCollectionSet)
			userCollection.DELETE("/:setID", handler.DeleteUserCollectionLegoSet)
		}
	}
}
