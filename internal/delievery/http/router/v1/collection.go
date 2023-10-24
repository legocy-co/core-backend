package v1

import (
	"github.com/gin-gonic/gin"
	a "legocy-go/internal/app"
	collection "legocy-go/internal/delievery/http/handlers/collection/user_collection"
	"legocy-go/internal/delievery/http/middleware"
)

func (router V1router) addUserCollections(rg *gin.RouterGroup, app *a.App) {
	handler := collection.NewUserLegoCollectionHandler(app.GetUserCollectionService())

	userCollection := rg.Group("/collections")
	{
		userCollection.GET("/", handler.GetUserCollection)
		userCollection.POST("/", handler.AddLegoSetToUserCollection)
		userCollection.GET("/valuation/:currencyID", handler.GetUserCollectionValuation)
		userCollection.Use(
			middleware.CollectionSetOwnerOrAdmin("setID", app.GetUserLegoSetsRepository()))
		{
			userCollection.PUT("/:setID", handler.UpdateUserCollectionSet)
			userCollection.DELETE("/:setID", handler.DeleteUserCollectionLegoSet)
		}
	}
}
