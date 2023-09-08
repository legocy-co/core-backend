package marketplace

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/delievery/http/resources/users"
	s "legocy-go/internal/domain/marketplace/service"
	ser "legocy-go/internal/domain/users/service"
	"legocy-go/pkg/eventNotifier/client"
	"net/http"
	"strconv"
)

type UserProfilePageHandler struct {
	marketItemService s.MarketItemService
	userService       ser.UserUseCase
	userReviewService s.UserReviewService
	userImageService  ser.UserImageUseCase
	notifyClient      client.EventNotifierClient
}

func NewUserProfilePageHandler(
	marketItemService s.MarketItemService, userService ser.UserUseCase,
	userReviewService s.UserReviewService, userImageService ser.UserImageUseCase,
	notifyClient client.EventNotifierClient) UserProfilePageHandler {

	return UserProfilePageHandler{
		marketItemService: marketItemService,
		userService:       userService,
		userReviewService: userReviewService,
		userImageService:  userImageService,
		notifyClient:      notifyClient,
	}
}

// UserProfilePageDetail
//
//	@Summary	Get User Profile Page
//	@Tags		user_profile_pages
//	@ID			detail_user_profile_page
//	@Param		userID	path	int	true	"user ID"
//	@Produce	json
//	@Success	200	{object}	marketplace.UserProfilePageResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/user-profile-pages/{userID} [get]
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

	user, err := h.userService.GetUserByID(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse := users.GetUserDetailResponse(user)

	userReviews, err := h.userReviewService.UserReviewsBySellerID(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userReviewsResponse := make([]marketplace.UserReviewResponse, 0, len(userReviews))
	for _, ur := range userReviews {
		userReviewsResponse = append(userReviewsResponse, marketplace.GetUserReviewResponse(ur))
	}

	userImages, err := h.userImageService.GetUserImages(c, userID)

	userImagesResponse := make([]users.UserImageInfoResponse, 0, len(userImages))
	for _, ui := range userImages {
		userImagesResponse = append(userImagesResponse, users.GetUserInfoResponse(ui))
	}

	userProfilePageResponse := marketplace.GetUserProfilePageResponse(
		marketItemsResponse, userResponse, userReviewsResponse, userImagesResponse)
	c.JSON(http.StatusOK, userProfilePageResponse)
}
