package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AddDefaultPagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := ctx.Query("page")
		limit := ctx.Query("page")

		if page == "" {
			page = "1"
			log.Println("Setting default page param: ", ctx.Param("page"))
		}
		if limit == "" {
			limit = "15"
			log.Println("Setting default limit param: ", ctx.Param("limit"))
		}

		ctx.AddParam("page", page)
		ctx.AddParam("limit", limit)
		ctx.Next()
	}
}
