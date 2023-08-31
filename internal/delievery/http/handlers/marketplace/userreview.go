package marketplace

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"legocy-go/internal/config"
	"legocy-go/internal/delievery/http/middleware"
	resources "legocy-go/internal/delievery/http/resources"
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/delievery/http/resources/pagination"
	s "legocy-go/internal/delievery/http/service/marketplace"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
	"legocy-go/internal/domain/users/middleware"
	"legocy-go/pkg/eventNotifier/client"
	clientModels "legocy-go/pkg/eventNotifier/models"
	"net/http"
	"strconv"
)

type UserReviewHandler struct {
	service      s.UserReviewService
	notifyClient client.EventNotifierClient
}

func NewUserReviewHandler(
	service s.UserReviewService, notifyClient client.EventNotifierClient) UserReviewHandler {

	return UserReviewHandler{
		service:      service,
		notifyClient: notifyClient,
	}
}

// ListUserReviews
//
//	@Summary	Get User Reviews
//	@Tags		user_reviews
//	@ID			list_user_reviews
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/user-reviews/ [get]
//
//	@Security	JWT
func (h *UserReviewHandler) ListUserReviews(c *gin.Context) {

	ctx := pagination.GetPaginationContext(c)

	var userReviews []*models.UserReview
	userReviews, err := h.service.ListUserReviews(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
	}

	if len(userReviews) == 0 {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"error": errors.ErrUserReviewsNotFound.Error()})
		return
	}

	userReviewResponse := make([]marketplace.UserReviewResponse, 0, len(userReviews))
	for _, m := range userReviews {
		userReviewResponse = append(userReviewResponse, marketplace.GetUserReviewResponse(m))
	}

	response := resources.DataMetaResponse{
		Data: userReviewResponse,
		Meta: pagination.GetPaginatedMetaResponse(
			c.Request.URL.Path, resources.MsgSuccess, ctx),
	}
	resources.Respond(c.Writer, response)
}

// UserReviewDetail
//
//		@Summary	Get User Review
//		@Tags		user_reviews
//		@ID			detail_user_review
//	 	@Param 		reviewID 	path  int true "review ID"
//		@Produce	json
//		@Success	200	{object}	marketplace.UserReviewResponse
//		@Failure	400	{object}	map[string]interface{}
//		@Router		/user-reviews/{reviewID} [get]
//
//		@Security	JWT
func (h *UserReviewHandler) UserReviewDetail(c *gin.Context) {
	reviewID, err := strconv.Atoi(c.Param("reviewID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	userReview, err := h.service.UserReviewDetail(c, reviewID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userReviewResponse := marketplace.GetUserReviewResponse(userReview)
	c.JSON(http.StatusOK, userReviewResponse)
}

// CreateUserReview
//
//		@Summary	Create User Review
//		@Tags		user_reviews
//		@ID			create_user_review
//	 	@Param 		data 	body marketplace.UserReviewRequest true "data"
//		@Produce	json
//		@Success	200	{object}	map[string]interface{}
//		@Failure	400	{object}	map[string]interface{}
//		@Router		/user-reviews/ [post]
//
//		@Security	JWT
func (h *UserReviewHandler) CreateUserReview(c *gin.Context) {
	// If we get here, then token payload is valid
	tokenString := v1.GetAuthTokenHeader(c)
	userPayload, ok := auth.ParseTokenClaims(tokenString)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "invalid token credentials"})
		return
	}

	var reviewRequest *marketplace.UserReviewRequest
	if err := c.ShouldBindJSON(&reviewRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Payload ID as ReviewerID
	userReviewValueObject, err := reviewRequest.ToUserReviewValueObject(userPayload.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateUserReview(c, userReviewValueObject)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Track Event
	h.notifyClient.NotifyEvent(clientModels.NotifyEventData{
		ChatID:  config.GetAppConfig().EventNotifierChatID,
		Message: fmt.Sprint("New UserReview created!"),
	})

	response := resources.DataMetaResponse{
		Data: reviewRequest,
		Meta: resources.SuccessMetaResponse,
	}
	resources.Respond(c.Writer, response)
}

// DeleteUserReview
//
//		@Summary	Delete User Review
//		@Tags		user_reviews
//		@ID			delete_user_review
//	 	@Param 		reviewId 	path	int true "review ID"
//		@Produce	json
//		@Success	200	{object}	map[string]bool
//		@Failure	400	{object}	map[string]interface{}
//		@Router		/user-reviews/{reviewId} [delete]
//
//		@Security	JWT
func (h *UserReviewHandler) DeleteUserReview(c *gin.Context) {
	reviewID, err := strconv.Atoi(c.Param("reviewId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	err = h.service.DeleteUserReview(c, reviewID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			"Error deleting UserReview object")
	}

	c.JSON(http.StatusOK, map[string]bool{"status": true})
}
