package fb

import "github.com/gin-gonic/gin"

func (h Handler) SignIn(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "TODO"})
}
