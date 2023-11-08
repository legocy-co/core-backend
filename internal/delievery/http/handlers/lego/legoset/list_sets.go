package legoset

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	"legocy-go/internal/delievery/http/errors"
	"legocy-go/internal/delievery/http/resources/lego"
	"legocy-go/internal/delievery/http/resources/pagination"
	"net/http"
)

// ListSets
//
//	@Summary	List of LegoSet objects
//	@Tags		lego_sets
//	@ID			list_sets
//	@Produce	json
//	@Success	200	{object}	[]lego.LegoSetResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/sets/ [get]
//
//	@Security	JWT
func (lsh *LegoSetHandler) ListSets(c *gin.Context) {

	ctx := pagination.GetPaginationContext(c)

	setsList, err := lsh.service.ListLegoSets(ctx)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	setsResponse := make([]lego.LegoSetResponse, 0, len(setsList))
	for _, legoSet := range setsList {
		setsResponse = append(setsResponse, lego.GetLegoSetResponse(legoSet))
	}

	c.JSON(http.StatusOK, setsResponse)
}
