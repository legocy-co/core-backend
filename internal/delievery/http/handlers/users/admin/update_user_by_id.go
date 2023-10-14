package admin

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/resources/users/admin"
	"net/http"
	"strconv"
)

// UpdateUserByID
//
//	@Summary	Update User
//	@Tags		users_admin
//	@ID			update_user
//	@Param		userId	path	int	true  "user ID"
//	@Param		data	body	admin.UserAdminUpdateRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	admin.UserAdminDetailResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/users/{userId} [put]
//
//	@Security	ApiKeyAuth
//	@param		Authorization	header	string	true	"Authorization"
func (uah *UserAdminHandler) UpdateUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	var userRequest *admin.UserAdminUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	vo, err := userRequest.ToUserAdminValueObject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	userDomain, err := uah.service.UpdateUser(c, userID, vo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse := admin.GetUserAdminDetailResponse(userDomain)
	c.JSON(http.StatusOK, userResponse)
}