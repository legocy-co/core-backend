package admin

import (
	"github.com/gin-gonic/gin"
	resources "legocy-go/internal/delievery/http/resources/users/admin"
	"net/http"
)

// AdminRegister
//
//	@Summary	Create Admin
//	@Tags		users_admin
//	@ID			create-admin
//	@Produce	json
//	@Param		data	body		resources.AdminRegistrationRequest	true	"reg request"
//	@Success	200		{object}	resources.AdminRegistrationResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/admin/users/register [post]
//
//	@Security	ApiKeyAuth
//	@param		Authorization	header	string	true	"Authorization"
func (uah *UserAdminHandler) AdminRegister(c *gin.Context) {

	var registerReq resources.AdminRegistrationRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAdmin := registerReq.ToAdmin()
	if err := uah.service.CreateAdmin(c, userAdmin, registerReq.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resources.GetAdminResponse(userAdmin))
}
