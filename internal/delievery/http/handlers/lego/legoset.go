package lego

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/resources"
	"legocy-go/internal/delievery/http/resources/lego"
	"legocy-go/internal/delievery/http/resources/pagination"
	s "legocy-go/internal/delievery/http/service/lego"
	"net/http"
	"strconv"
)

type LegoSetHandler struct {
	service s.LegoSetUseCase
}

func NewLegoSetHandler(service s.LegoSetUseCase) LegoSetHandler {
	return LegoSetHandler{service: service}
}

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

// SetDetail
//
//	@Summary	Get LegoSet by ID
//	@Tags		lego_sets
//	@ID			set_detail
//	@Param		setID	path	int	true	"Lego Set ID"
//	@Produce	json
//	@Success	200	{object}	lego.LegoSetResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/sets/{setID} [get]
//
//	@Security	JWT
func (lsh *LegoSetHandler) SetDetail(c *gin.Context) {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	legoSet, err := lsh.service.LegoSetDetail(c.Request.Context(), setID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	legoSetResponse := lego.GetLegoSetResponse(legoSet)

	c.JSON(http.StatusOK, legoSetResponse)
}

// SetCreate
//
//	@Summary	Create Lego Set object
//	@Tags		lego_sets_admin
//	@ID			set_create
//	@Param		data	body	lego.LegoSetRequest	true	"create data"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/sets/ [post]
//
//	@Security	JWT
func (lsh *LegoSetHandler) SetCreate(c *gin.Context) {
	var setRequest lego.LegoSetRequest
	if err := c.ShouldBindJSON(&setRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	legoSetBasic := setRequest.ToLegoSeriesBasic()
	err := lsh.service.LegoSetCreate(c.Request.Context(), legoSetBasic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	response := v1.DataMetaResponse{
		Data: true,
		Meta: v1.SuccessMetaResponse,
	}

	v1.Respond(c.Writer, response)
}

// SetDelete
//
//	@Summary	Delete Lego Set object
//	@Tags		lego_sets_admin
//	@ID			set_delete
//	@Param		setID	path	int	true	"LegoSet ID"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/sets/{setID} [delete]
//
//	@Security	JWT
func (lsh *LegoSetHandler) SetDelete(c *gin.Context) {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	err = lsh.service.LegoSetDelete(c, setID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	response := v1.DataMetaResponse{
		Data: true,
		Meta: v1.SuccessMetaResponse,
	}

	v1.Respond(c.Writer, response)
}
