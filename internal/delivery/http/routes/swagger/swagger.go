package swagger

import (
	"github.com/gin-gonic/gin"
	_ "github.com/legocy-co/legocy/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddSwaggerDocs(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
