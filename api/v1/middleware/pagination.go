package v1

import "github.com/gin-gonic/gin"

func AddDefaultPagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if page := ctx.Query("page"); page == "" {
			ctx.Request.URL.Query().Set("page", "1")
		}
		if limit := ctx.Query("limit"); limit == "" {
			ctx.Request.URL.Query().Set("limit", "15")
		}
	}
}
