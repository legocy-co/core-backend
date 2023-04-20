package auth

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	resources "legocy-go/internal/delievery/http/resources/users"
	jwt "legocy-go/internal/domain/users/middleware"
	"net/http"
)

// GenerateToken
//
//	@Summary	generate jwt token
//	@Tags		authentication
//	@ID			create-jwt
//	@Produce	json
//	@Param		data	body		resources.JWTRequest	true	"jwt request"
//	@Success	200		{object}	resources.JWTResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/token [post]
func (th *TokenHandler) GenerateToken(c *gin.Context) {

	var jwtRequest resources.JWTRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	err := th.service.ValidateUser(c.Request.Context(), jwtRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := th.service.GetUserByEmail(c.Request.Context(), jwtRequest.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateJWT(user.Email, user.ID, user.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resources.JWTResponse{
		AccessToken: token,
	})

}
