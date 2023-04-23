package legoset

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	v1 "legocy-go/internal/delievery/http/resources"
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
		v1.ErrorRespond(c.Writer, err.Error())
		return
	}

	setsResponse := make([]lego.LegoSetResponse, 0, len(setsList))
	for _, legoSet := range setsList {
		setsResponse = append(setsResponse, lego.GetLegoSetResponse(legoSet))
	}

	if len(setsResponse) == 0 {
		v1.ErrorRespond(c.Writer, "No data found")
		return
	}

	c.JSON(http.StatusOK, setsResponse)
}
