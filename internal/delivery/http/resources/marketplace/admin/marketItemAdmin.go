package admin

import (
	"legocy-go/internal/delivery/http/resources/lego"
	"legocy-go/internal/delivery/http/resources/users"
	legoDomain "legocy-go/internal/domain/lego/models"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemAdminCreateRequest struct {
	LegoSetID   int     `json:"lego_set_id"`
	SellerID    int     `json:"seller_id"`
	Price       float32 `json:"price"`
	Location    string  `json:"location"`
	Status      string  `json:"status"`
	SetState    string  `json:"set_state"`
	Description string  `json:"description"`
}

func (r MarketItemAdminCreateRequest) ToMarketItemAdminValueObject() (
	*models.MarketItemAdminValueObject, error) {

	if !models.IsValidListingStatus(r.Status) {
		return nil, errors.ErrMarketItemInvalidStatus
	}

	if !legoDomain.IsValidSetState(r.SetState) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.MarketItemAdminValueObject{
		LegoSetID:   r.LegoSetID,
		SellerID:    r.SellerID,
		Price:       r.Price,
		Location:    r.Location,
		Status:      r.Status,
		SetState:    r.SetState,
		Description: r.Description,
	}, nil
}

type MarketItemAdminUpdateRequest struct {
	LegoSetID   int     `json:"lego_set_id"`
	SellerID    int     `json:"seller_id"`
	Price       float32 `json:"price"`
	Location    string  `json:"location"`
	Status      string  `json:"status"`
	SetState    string  `json:"set_state"`
	Description string  `json:"description"`
}

func (r MarketItemAdminUpdateRequest) ToMarketItemAdminValueObject() (
	*models.MarketItemAdminValueObject, error) {

	if !models.IsValidListingStatus(r.Status) {
		return nil, errors.ErrMarketItemInvalidStatus
	}

	if !legoDomain.IsValidSetState(r.SetState) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.MarketItemAdminValueObject{
		LegoSetID:   r.LegoSetID,
		SellerID:    r.SellerID,
		Price:       r.Price,
		Location:    r.Location,
		Status:      r.Status,
		SetState:    r.SetState,
		Description: r.Description,
	}, nil
}

type MarketItemAdminResponse struct {
	ID          int                      `json:"id"`
	Price       float32                  `json:"price"`
	Location    string                   `json:"location"`
	LegoSet     lego.LegoSetResponse     `json:"lego_set"`
	Seller      users.UserDetailResponse `json:"seller"`
	Status      string                   `json:"status"`
	SetState    string                   `json:"set_state"`
	Description string                   `json:"description"`
}

func GetMarketItemAdminResponse(
	mi *models.MarketItemAdmin) MarketItemAdminResponse {
	return MarketItemAdminResponse{
		ID:          mi.ID,
		Price:       mi.Price,
		Location:    mi.Location,
		LegoSet:     lego.GetLegoSetResponse(&mi.LegoSet),
		Seller:      users.GetUserDetailResponse(&mi.Seller),
		Status:      mi.Status,
		SetState:    mi.SetState,
		Description: mi.Description,
	}
}
