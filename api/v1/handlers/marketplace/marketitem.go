package marketplace

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "legocy-go/api/v1/middleware"
	r "legocy-go/api/v1/resources"
	res "legocy-go/api/v1/resources/marketplace"
	s "legocy-go/api/v1/usecase/marketplace"
	"legocy-go/internal/utils"
	auth "legocy-go/pkg/auth/middleware"
	"legocy-go/pkg/marketplace/errors"
	models "legocy-go/pkg/marketplace/models"
	"net/http"
	"strconv"
)

type MarketItemHandler struct {
	service s.MarketItemService
}

func NewMarketItemHandler(service s.MarketItemService) MarketItemHandler {
	return MarketItemHandler{
		service: service,
	}
}

func (h *MarketItemHandler) ListMarketItems(c *gin.Context) {
	var marketItems []*models.MarketItem
	var err error

	var pagination = utils.Pagination{
		Page:  c.Param("page"),
		Limit: c.Param("limit"),
	}

	b := context.Background()
	ctx := context.WithValue(b, "pagination", pagination)

	sellerId := c.Query("sellerId")
	if sellerId == "" {
		// Load all
		marketItems, err = h.service.ListMarketItems(ctx)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		// Load seller MarketItems
		sellerId, err := strconv.Atoi(sellerId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		marketItems, err = h.service.SellerMarketItems(ctx, sellerId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if len(marketItems) == 0 {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": errors.ErrMarketItemsNotFound})
		return
	}

	var marketItemResponse []res.MarketItemResponse
	for _, m := range marketItems {
		marketItemResponse = append(marketItemResponse, res.GetMarketItemResponse(m))
	}

	response := r.DataMetaResponse{
		Data: marketItemResponse,
		Meta: r.GetPaginatedMetaResponse(c.Request.URL.Path, r.MsgSuccess, pagination),
	}
	r.Respond(c.Writer, response)
}

func (h *MarketItemHandler) MarketItemDetail(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("itemID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	marketItem, err := h.service.MarketItemDetail(c, itemID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marketItemResponse := res.GetMarketItemResponse(marketItem)
	response := r.DataMetaResponse{
		Data: marketItemResponse,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
}

func (h *MarketItemHandler) CreateMarketItem(c *gin.Context) {
	// If we get here, then token payload is valid
	tokenString := v1.GetAuthTokenHeader(c)
	userPayload, ok := auth.ParseTokenClaims(tokenString)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "invalid token credentials"})
		return
	}

	var itemRequest *res.MarketItemRequest
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Payload ID as SellerID
	err := h.service.CreateMarketItem(c, itemRequest.ToMarketItemBasic(userPayload.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := r.DataMetaResponse{
		Data: itemRequest,
		Meta: r.SuccessMetaResponse,
	}
	r.Respond(c.Writer, response)
}

func (h *MarketItemHandler) DeleteMarketItem(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	err = h.service.DeleteMarketItem(c, itemID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			"Error deleting MarketItem object")
	}

	response := r.DataMetaResponse{
		Data: true,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
}
