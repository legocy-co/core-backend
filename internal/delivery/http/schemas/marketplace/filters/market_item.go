package filters

import (
	"github.com/gin-gonic/gin"
	legoFilters "github.com/legocy-co/legocy/internal/delivery/http/schemas/lego/filters"
	legoFilterDomain "github.com/legocy-co/legocy/internal/domain/lego/filters"
	domain "github.com/legocy-co/legocy/internal/domain/marketplace/filters"
	"github.com/legocy-co/legocy/internal/pkg/errors"
	"github.com/legocy-co/legocy/lib/helpers"
)

func GetMarketItemFilterCritera(ctx *gin.Context) (*domain.MarketItemFilterCriteria, *errors.AppError) {
	var filterDTO MarketItemFilterDTO

	if err := helpers.BindQueryParamsToStruct(&filterDTO, ctx.Request.URL.Query()); err != nil {
		appErr := errors.NewAppError(errors.ValidationError, err.Error())
		return nil, &appErr
	}

	return filterDTO.ToCriteria()
}

type MarketItemFilterDTO struct {
	SetIDs        []int                         `form:"set_id__in" json:"set_id__in"`
	PriceGTE      *float64                      `form:"price_gte" json:"price_gte"`
	PriceLTE      *float64                      `form:"price_lte" json:"price_lte"`
	SetStates     []string                      `form:"set_state__in" json:"set_state__in"`
	Locations     []string                      `form:"location__in" json:"location__in"`
	LegoSet       *legoFilters.LegoSetFilterDTO `form:"lego_set" json:"lego_set"`
	MarketItemIds string                        `form:"market_item_ids" json:"market_item_ids"`
}

func (dto *MarketItemFilterDTO) ToCriteria() (*domain.MarketItemFilterCriteria, *errors.AppError) {

	var legoSetFilter *legoFilterDomain.LegoSetFilterCriteria = nil
	if dto.LegoSet != nil {
		legoSetFilter = dto.LegoSet.ToCriteria()
	}

	ids := make([]int, 0)
	if dto.MarketItemIds != "" {
		ids = helpers.StringToIntArray(dto.MarketItemIds, ",")
	}

	return domain.NewMarketItemFilterCriteria(
		ids,
		dto.SetIDs,
		dto.PriceGTE,
		dto.PriceLTE,
		dto.SetStates,
		dto.Locations,
		legoSetFilter,
	)
}
