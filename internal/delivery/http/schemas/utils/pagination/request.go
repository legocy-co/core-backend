package pagination

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/pkg/pagination"
	"strconv"
)

func GetPaginationContext(c *gin.Context) pagination.PaginationContext {

	var limit, offset int

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err = strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	return pagination.NewPaginationContext(limit, offset)
}
