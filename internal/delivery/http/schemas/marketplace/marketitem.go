package marketplace

import (
	legoResources "github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	lego "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/domain/marketplace/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
)

type MarketItemRequest struct {
	LegoSetID   int     `json:"legoSetID"`
	Price       float32 `json:"price"`
	Location    string  `json:"location"`
	SetState    string  `json:"setState"`
	Description string  `json:"description"`
	IsSold      bool    `json:"isSold"`
	Changed     bool    `json:"changed"`
}

func (r *MarketItemRequest) ToMarketItemValueObject(sellerID int) (*models.MarketItemValueObject, error) {

	if !lego.IsValidSetState(r.SetState) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	var status = models.ListingStatusActive
	if r.IsSold {
		status = models.ListingStatusSold
	} else if r.Changed {
		status = models.ListingStatusCheckRequired
	}

	return &models.MarketItemValueObject{
		LegoSetID:   r.LegoSetID,
		SellerID:    sellerID,
		Price:       r.Price,
		Location:    r.Location,
		Status:      status,
		Description: r.Description,
		SetState:    r.SetState,
	}, nil
}

type MarketItemResponse struct {
	ID          int                           `json:"id"`
	Price       float32                       `json:"price"`
	Location    string                        `json:"location"`
	LegoSet     legoResources.LegoSetResponse `json:"legoSet"`
	Seller      users.UserDetailResponse      `json:"seller"`
	Status      string                        `json:"status"`
	SetState    string                        `json:"setState"`
	Description string                        `json:"description"`
	Images      []ImageResponse               `json:"images"`
	IsLiked     bool                          `json:"isLiked"`
}

func GetMarketItemResponse(m *models.MarketItem) MarketItemResponse {
	return MarketItemResponse{
		ID:          m.ID,
		Price:       m.Price,
		Location:    m.Location,
		LegoSet:     legoResources.GetLegoSetResponse(&m.LegoSet),
		Seller:      users.GetUserDetailResponse(&m.Seller),
		Status:      m.Status,
		SetState:    m.SetState,
		Description: m.Description,
		Images:      GetImagesResponse(m.Images),
		IsLiked:     m.Liked,
	}
}

func (r *MarketItemResponse) WithReviewsTotals(totals *users.UserReviewTotalsResponse) *MarketItemResponse {
	r.Seller = *r.Seller.WithReviewTotals(totals)
	return r
}

func GetMarketItemResponseWithReviewTotals(
	m *models.MarketItem,
	totals *models.UserRevewTotals) MarketItemResponse {

	r := GetMarketItemResponse(m)

	if totals == nil {
		return r
	}

	totalsResponse := users.GetUserReviewsTotalsResponse(totals)
	return *r.WithReviewsTotals(totalsResponse)
}
