package v1

import (
	res "legocy-go/api/v1/resources"
	s "legocy-go/api/v1/usecase"

	"github.com/gin-gonic/gin"
)

type LegoSeriesHandler struct {
	service s.LegoSeriesService
}

func NewLegoSeriesHandler(service s.LegoSeriesService) LegoSeriesHandler {
	return LegoSeriesHandler{service: service}
}

func (lsh *LegoSeriesHandler) ListSeries(c *gin.Context) {
	seriesList, err := lsh.service.ListSeries(c.Request.Context())
	if err != nil {
		res.ErrorRespond(c.Writer, "Error extracting LEGO Series List")
	}

	seriesResponse := res.DataMetaResponse{
		Data: seriesList,
		Meta: map[string]interface{}{
			"status": 200,
			"msg":    "OK",
		},
	}

	res.Respond(c.Writer, seriesResponse)
}
