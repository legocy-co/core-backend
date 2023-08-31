package auth

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	resources "legocy-go/internal/delievery/http/resources/users"
	"net/http"
)

// AdminRegister
//
//	@Summary	Create Admin User
//	@Tags		authentication_admin
//	@ID			create-admin
//	@Produce	json
//	@Param		data	body		resources.UserRegistrationRequest	true	"reg request"
//	@Success	200		{object}	resources.UserRegistrationResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/admin/users [post]
//
//	@Security	ApiKeyAuth
//	@param		Authorization	header	string	true	"Authorization"
func (th *TokenHandler) AdminRegister(c *gin.Context) {

	var registerReq resources.UserRegistrationRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := registerReq.ToAdmin()
	if err := th.service.CreateUser(c, user, registerReq.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resources.GetUserResponse(user))
}
