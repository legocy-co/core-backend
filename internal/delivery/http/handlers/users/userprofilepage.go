package users

import (
	"github.com/gin-gonic/gin"
	appErrors "github.com/legocy-co/legocy/internal/app/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users/profile"
	s "github.com/legocy-co/legocy/internal/domain/marketplace/service"
	ser "github.com/legocy-co/legocy/internal/domain/users/service"
	"net/http"
	"strconv"
)

type UserProfilePageHandler struct {
	marketItemService s.MarketItemService
	userService       ser.UserService
	userReviewService s.UserReviewService
}

func NewUserProfilePageHandler(
	marketItemService s.MarketItemService,
	userService ser.UserService,
	userReviewService s.UserReviewService) UserProfilePageHandler {

	return UserProfilePageHandler{
		marketItemService: marketItemService,
		userService:       userService,
		userReviewService: userReviewService,
	}
}

// UserProfilePageDetail
//
//	@Summary	Get User Profile Page
//	@Tags		user_profile_pages
//	@ID			detail_user_profile_page
//	@Param		userID	path	int	true	"user ID"
//	@Produce	json
//	@Success	200	{object}	profile.UserProfilePageResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/users/profile/{userID} [get]
//
//	@Security	JWT
func (h *UserProfilePageHandler) UserProfilePageDetail(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	marketItems, err := h.marketItemService.MarketItemsBySellerID(c, userID)

	marketItemsResponse := make([]marketplace.MarketItemResponse, 0, len(marketItems))
	for _, mi := range marketItems {
		marketItemsResponse = append(marketItemsResponse, marketplace.GetMarketItemResponse(mi))
	}

	user, appErr := h.userService.GetUserByID(c, userID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	userResponse := users.GetUserDetailResponse(user)

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
