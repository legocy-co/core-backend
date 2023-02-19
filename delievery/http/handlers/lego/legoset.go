package lego

import (
	"github.com/gin-gonic/gin"
	r "legocy-go/delievery/http/resources"
	res "legocy-go/delievery/http/resources/lego"
	s "legocy-go/delievery/http/usecase/lego"
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
		r.ErrorRespond(c.Writer, err.Error())
		return
	}

	setsResponse := make([]res.LegoSetResponse, 0, len(setsList))
	for _, legoSet := range setsList {
		setsResponse = append(setsResponse, res.GetLegoSetResponse(legoSet))
	}

	if len(setsResponse) == 0 {
		r.ErrorRespond(c.Writer, "No data found")
		return
	}

	response := r.DataMetaResponse{
		Data: setsResponse,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
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

	response := r.DataMetaResponse{
		Data: legoSetResponse,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
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

	response := r.DataMetaResponse{
		Data: true,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
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

	response := r.DataMetaResponse{
		Data: true,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
}
