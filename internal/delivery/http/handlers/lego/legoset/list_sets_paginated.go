package legoset

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
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
//
// @Param       filter  query  	filters.LegoSetFilterDTO false "filter"
//
//	@Success	200	{object}	pagination.PageResponse[lego.LegoSetResponse]
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/sets/ [get]
//
// @Security JWT
func (h *LegoSetHandler) ListSetsPaginated(c *gin.Context) {

	ctx := pagination.GetPaginationContext(c)
	filterCriteria := filters.GetLegoSetFilterCritera(c)

	setsPage, appErr := h.service.GetSetsPage(ctx, filterCriteria)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
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
