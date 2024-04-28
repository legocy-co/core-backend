package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users/profile"
	appErrors "github.com/legocy-co/legocy/internal/pkg/app/errors"
	"net/http"
	"strconv"
)

// CurrentUserProfilePage
//
//	@Summary	Get User Profile Page
//	@Tags		user_profile_pages
//	@ID			current_user_profile_page
//	@Param		userID	path	int	true	"user ID"
//	@Produce	json
//	@Success	200	{object}	profile.UserProfilePageResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/users/profile/ [get]
//
//	@Security	JWT
func (h *UserProfilePageHandler) CurrentUserProfilePage(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	marketItems, err := h.marketItemService.GetMarketItemsBySellerID(c, userID)

	// Get Active MarketItems for User
	marketItemsResponse := make([]marketplace.MarketItemResponse, 0, len(marketItems))
	for _, mi := range marketItems {
		marketItemsResponse = append(marketItemsResponse, marketplace.GetMarketItemResponse(mi))
	}

	// Get User
	user, appErr := h.userService.GetUserByID(c, userID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	// Get Review Totals
	userReviewsTotals, appErr := h.userReviewService.GetUserReviewsTotals(c, userID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	reviewsTotalsResponse := users.GetUserReviewsTotalsResponse(userReviewsTotals)

	userResponse := users.GetUserDetailResponse(user)
	userResponse = *userResponse.WithReviewTotals(reviewsTotalsResponse)

	// Get User Reviews
	userReviews, appErr := h.userReviewService.UserReviewsBySellerID(c, userID)
	if appErr != nil && appErr.GetErrorType() != appErrors.NotFoundError {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	userReviewsResponse := make([]users.UserReviewResponse, 0, len(userReviews))
	for _, ur := range userReviews {
		userReviewsResponse = append(userReviewsResponse, users.GetUserReviewResponse(ur))
	}

	userProfilePageResponse := profile.GetUserProfilePageResponse(
		marketItemsResponse, userResponse, userReviewsResponse)
	c.JSON(http.StatusOK, userProfilePageResponse)
}
