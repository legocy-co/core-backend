package filters

import (
	domain "github.com/legocy-co/legocy/internal/domain/lego/filters"
)

type LegoSetFilterDTO struct {
	NpiecesGTE *int      `form:"npieces_gte"`
	NpiecesLTE *int      `form:"npieces_lte"`
	SeriesIDs  *[]int    `form:"series_id__in"`
	SetNumbers *[]string `form:"set_number__in"`
	Name       *string   `form:"name__icontains"`
}

func (dto LegoSetFilterDTO) ToCriteria() *domain.LegoSetFilterCriteria {
	return domain.NewLegoSetFilterCriteria(
		dto.NpiecesGTE,
		dto.NpiecesLTE,
		dto.SeriesIDs,
		dto.SetNumbers,
		dto.Name,
	)
}
