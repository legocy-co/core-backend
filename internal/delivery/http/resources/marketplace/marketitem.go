package marketplace

import (
	legoResources "legocy-go/internal/delivery/http/resources/lego"
	"legocy-go/internal/delivery/http/resources/users"
	lego "legocy-go/internal/domain/lego/models"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemRequest struct {
	LegoSetID   int     `json:"lego_set_id"`
	Price       float32 `json:"price"`
	Location    string  `json:"location"`
	SetState    string  `json:"set_state"`
	Description string  `json:"description"`
}

func (r *MarketItemRequest) ToMarketItemValueObject(sellerID int) (*models.MarketItemValueObject, error) {

	if !lego.IsValidSetState(r.SetState) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.MarketItemValueObject{
		LegoSetID:   r.LegoSetID,
		SellerID:    sellerID,
		Price:       r.Price,
		Location:    r.Location,
		Status:      models.ListingStatusCheckRequired,
		Description: r.Description,
	}, nil
}

type MarketItemResponse struct {
	ID          int                           `json:"id"`
	Price       float32                       `json:"price"`
	Location    string                        `json:"location"`
	LegoSet     legoResources.LegoSetResponse `json:"lego_set"`
	Seller      users.UserDetailResponse      `json:"seller"`
	Status      string                        `json:"status"`
	SetState    string                        `json:"set_state"`
	Description string                        `json:"description"`
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
	}
}
