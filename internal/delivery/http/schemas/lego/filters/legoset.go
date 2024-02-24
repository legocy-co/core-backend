package filters

import (
	domain "github.com/legocy-co/legocy/internal/domain/lego/filters"
)

type LegoSetFilterDTO struct {
	NpiecesGTE *int      `form:"npieces_gte" json:"npieces_gte"`
	NpiecesLTE *int      `form:"npieces_lte" json:"npieces_lte"`
	SeriesIDs  *[]int    `form:"series_id__in" json:"series_id__in"`
	SetNumbers *[]string `form:"set_number__in" json:"set_number__in"`
	Name       *string   `form:"name__ilike" json:"name__ilike"`
}

func (dto *LegoSetFilterDTO) ToCriteria() *domain.LegoSetFilterCriteria {

	if dto == nil {
		return nil
	}

	return domain.NewLegoSetFilterCriteria(
		dto.NpiecesGTE,
		dto.NpiecesLTE,
		dto.SeriesIDs,
		dto.SetNumbers,
		dto.Name,
	)
}
