package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/legocy-co/legocy/docs"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"net/http"
)

// UserRegister
//
//	@Summary	register new user
//	@Tags		authentication
//	@ID			user-register
//	@Produce	json
//	@Param		data	body		schemas.UserRegistrationRequest	true	"user data"
//	@Success	200		{object}	schemas.UserRegistrationResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/auth/register [post]
func (th *TokenHandler) UserRegister(c *gin.Context) {

	var registerReq schemas.UserRegistrationRequest

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

	c.JSON(http.StatusOK, schemas.GetUserResponse(user))
}
