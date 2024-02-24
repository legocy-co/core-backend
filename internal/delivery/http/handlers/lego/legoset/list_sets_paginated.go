package legoset

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/lego/filters"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/utils/pagination"
	"net/http"
)

// ListSetsPaginated
//
//	@Summary	Get Lego Sets
//	@Tags		lego_sets
//	@ID			list_lego_sets_paginated
//	@Produce	json
//	@Param		limit	query	int	false	"limit" 10
//	@Param		offset	query	int	false	"offset" 0
//	@Success	200	{object}	pagination.PageResponse[lego.LegoSetResponse]
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/sets/ [get]
//
// @Security JWT
func (h *LegoSetHandler) ListSetsPaginated(c *gin.Context) {
	ctx := pagination.GetPaginationContext(c)

	var filterDTO *filters.LegoSetFilterDTO
	if err := c.ShouldBindQuery(&filterDTO); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	setsPage, err := h.service.GetSetsPage(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	legoSetResponse := make([]lego.LegoSetResponse, 0, len(setsPage.GetObjects()))
	for _, m := range setsPage.GetObjects() {
		legoSetResponse = append(legoSetResponse, lego.GetLegoSetResponse(&m))
	}

	responsePage := pagination.GetPageResponse[lego.LegoSetResponse](
		legoSetResponse,
		setsPage.GetTotal(),
		setsPage.GetLimit(),
		setsPage.GetOffset(),
	)

	c.JSON(http.StatusOK, responsePage)
}
