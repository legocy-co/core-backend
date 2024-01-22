package users

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/config"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	s "github.com/legocy-co/legocy/internal/domain/marketplace/service"
	"github.com/legocy-co/legocy/pkg/auth/jwt"
	"github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
	"net/http"
	"strconv"
)

type UserReviewHandler struct {
	service s.UserReviewService
}

func NewUserReviewHandler(
	service s.UserReviewService) UserReviewHandler {

	return UserReviewHandler{
		service: service,
	}
}

// ListUserReviews
//
//	@Summary	Get User Reviews
//	@Tags		user_reviews
//	@ID			list_user_reviews
//	@Produce	json
//	@Success	200	{object}	[]users.UserReviewResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/users/reviews/ [get]
//
//	@Security	JWT
func (h *UserReviewHandler) ListUserReviews(c *gin.Context) {

	var userReviews []*models.UserReview
	userReviews, err := h.service.ListUserReviews(c)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	userReviewResponse := make([]users.UserReviewResponse, 0, len(userReviews))
	for _, m := range userReviews {
		userReviewResponse = append(userReviewResponse, users.GetUserReviewResponse(m))
	}

	c.JSON(http.StatusOK, userReviewResponse)
}

// UserReviewDetail
//
//	@Summary	Get User Review
//	@Tags		user_reviews
//	@ID			detail_user_review
//	@Param		reviewID	path	int	true	"review ID"
//	@Produce	json
//	@Success	200	{object}	users.UserReviewResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/users/reviews/{reviewID} [get]
//
//	@Security	JWT
func (h *UserReviewHandler) UserReviewDetail(c *gin.Context) {
	reviewID, err := strconv.Atoi(c.Param("reviewID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	userReview, appErr := h.service.UserReviewDetail(c, reviewID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	userReviewResponse := users.GetUserReviewResponse(userReview)
	c.JSON(http.StatusOK, userReviewResponse)
}

// CreateUserReview
//
//	@Summary	Create User Review
//	@Tags		user_reviews
//	@ID			create_user_review
//	@Param		data	body	users.UserReviewRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/users/reviews/ [post]
//
//	@Security	JWT
func (h *UserReviewHandler) CreateUserReview(c *gin.Context) {
	// If we get here, then token payload is valid
	tokenString := middleware.GetAuthTokenHeader(c)
	userPayload, ok := jwt.ParseTokenClaims(tokenString, config.GetAppConfig().JwtConf.SecretKey)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "invalid token credentials"})
		return
	}

	var reviewRequest *users.UserReviewRequest
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

	appErr := h.service.CreateUserReview(c, userReviewValueObject)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// DeleteUserReview
//
//	@Summary	Delete User Review
//	@Tags		user_reviews
//	@ID			delete_user_review
//	@Param		reviewId	path	int	true	"review ID"
//	@Produce	json
//	@Success	200	{object}	map[string]bool
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/users/reviews/{reviewId} [delete]
//
//	@Security	JWT
func (h *UserReviewHandler) DeleteUserReview(c *gin.Context) {
	reviewID, err := strconv.Atoi(c.Param("reviewId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	appErr := h.service.DeleteUserReview(c, reviewID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
