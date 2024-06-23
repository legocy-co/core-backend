package filters

import (
	"github.com/gin-gonic/gin"
	domain "github.com/legocy-co/legocy/internal/domain/lego/filters"
	"github.com/legocy-co/legocy/lib/helpers"
)

func GetLegoSetFilterCriteria(ctx *gin.Context) *domain.LegoSetFilterCriteria {
	var filterDTO LegoSetFilterDTO

	if err := ctx.BindQuery(&filterDTO); err != nil {
		return nil
	}

	return filterDTO.ToCriteria()
}

type LegoSetFilterDTO struct {
	NpiecesGTE   *int    `form:"npieces_gte" json:"npieces_gte"`
	NpiecesLTE   *int    `form:"npieces_lte" json:"npieces_lte"`
	SeriesIDs    []int   `form:"series_id__in" json:"series_id__in"`
	SetNumbers   []int   `form:"set_number__in" json:"set_number__in"`
	Name         *string `form:"name__ilike" json:"name__ilike"`
	ReleaseYears string  `form:"release_year__in" json:"release_year__in"`
}

func (dto *LegoSetFilterDTO) ToCriteria() *domain.LegoSetFilterCriteria {

	if dto == nil {
		return nil
	}

	releaseYears := make([]int, 0)
	if dto.ReleaseYears != "" {
		releaseYears = helpers.StringToIntArray(dto.ReleaseYears, ",")
	}

	return domain.NewLegoSetFilterCriteria(
		dto.NpiecesGTE,
		dto.NpiecesLTE,
		dto.SeriesIDs,
		dto.SetNumbers,
		dto.Name,
		releaseYears,
	)
}
