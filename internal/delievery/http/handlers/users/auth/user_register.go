package auth

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	resources "legocy-go/internal/delievery/http/resources/users"
	"net/http"
)

// UserRegister
//
//	@Summary	register new user
//	@Tags		authentication
//	@ID			user-register
//	@Produce	json
//	@Param		data	body		resources.UserRegistrationRequest	true	"user data"
//	@Success	200		{object}	resources.UserRegistrationResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/auth/register [post]
func (th *TokenHandler) UserRegister(c *gin.Context) {

	var registerReq resources.UserRegistrationRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := registerReq.ToUser()
	if err := th.service.CreateUser(c, user, registerReq.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resources.GetUserResponse(user))
}
