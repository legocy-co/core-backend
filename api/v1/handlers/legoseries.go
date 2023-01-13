package v1

import (
	res "legocy-go/api/v1/resources"
	service "legocy-go/api/v1/usecase"

	"github.com/gin-gonic/gin"
)

func ListSeries(c *gin.Context, s service.LegoSeriesService) {
	seriesList, err := s.ListSeries(c.Request.Context())
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
