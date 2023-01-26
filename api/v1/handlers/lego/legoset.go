package lego

import (
	"github.com/gin-gonic/gin"
	res "legocy-go/api/v1/resources"
	ser "legocy-go/api/v1/usecase/lego"
	"net/http"
	"strconv"
)

type LegoSetHandler struct {
	service ser.LegoSetUseCase
}

func NewLegoSetHandler(service ser.LegoSetUseCase) LegoSetHandler {
	return LegoSetHandler{service: service}
}

func (lsh *LegoSetHandler) ListSets(c *gin.Context) {
	setsList, err := lsh.service.ListLegoSets(c)
	if err != nil {
		res.ErrorRespond(c.Writer, err.Error())
		return
	}

	setsResponse := res.DataMetaResponse{
		Data: setsList,
		Meta: res.SuccessMetaResponse,
	}

	res.Respond(c.Writer, setsResponse)
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

	legoSetResponse := res.DataMetaResponse{
		Data: legoSet,
		Meta: res.SuccessMetaResponse,
	}

	res.Respond(c.Writer, legoSetResponse)
}
