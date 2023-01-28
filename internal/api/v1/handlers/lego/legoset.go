package lego

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/api/v1/resources"
	res "legocy-go/internal/api/v1/resources/lego"
	s "legocy-go/internal/api/v1/usecase/lego"
	"net/http"
	"strconv"
)

type LegoSetHandler struct {
	service s.LegoSetUseCase
}

func NewLegoSetHandler(service s.LegoSetUseCase) LegoSetHandler {
	return LegoSetHandler{service: service}
}

func (lsh *LegoSetHandler) ListSets(c *gin.Context) {
	setsList, err := lsh.service.ListLegoSets(c)
	if err != nil {
		v1.ErrorRespond(c.Writer, err.Error())
		return
	}

	var setsResponse []res.LegoSetResponse
	for _, legoSet := range setsList {
		setsResponse = append(setsResponse, res.GetLegoSetResponse(legoSet))
	}

	if len(setsResponse) == 0 {
		v1.ErrorRespond(c.Writer, "No data found")
		return
	}

	response := v1.DataMetaResponse{
		Data: setsResponse,
		Meta: v1.SuccessMetaResponse,
	}

	v1.Respond(c.Writer, response)
}

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

	legoSetResponse := res.GetLegoSetResponse(legoSet)

	response := v1.DataMetaResponse{
		Data: legoSetResponse,
		Meta: v1.SuccessMetaResponse,
	}

	v1.Respond(c.Writer, response)
}

func (lsh *LegoSetHandler) SetCreate(c *gin.Context) {
	var setRequest res.LegoSetRequest
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

func (lsh *LegoSetHandler) SetDelete(c *gin.Context) {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	err = lsh.service.LegoSetDelete(c, setID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	response := v1.DataMetaResponse{
		Data: true,
		Meta: v1.SuccessMetaResponse,
	}

	v1.Respond(c.Writer, response)
}
