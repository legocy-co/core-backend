package v1

import (
	res "legocy-go/api/v1/resources"
	s "legocy-go/api/v1/usecase"
	jwt "legocy-go/pkg/auth/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context, service s.JWTUseCase) {

	var jwtRequest res.JWTRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	errIsValidUser := service.ValidateUser(&jwtRequest)
	if errIsValidUser != nil {
		res.ErrorRespond(c.Writer, "Invalid credentials")
		return
	}

	token, err := jwt.GenerateJWT(jwtRequest.Username, jwtRequest.Password)
	if err != nil {
		res.ErrorRespond(c.Writer, "Error generating token")
		return
	}

	res.Respond(c.Writer, res.DataMetaResponse{
		Data: res.JWTResponse{
			AccessToken: token,
		},
		Meta: map[string]interface{}{
			"msg":    res.MSG_SUCCESS,
			"status": 200,
		},
	})

}
