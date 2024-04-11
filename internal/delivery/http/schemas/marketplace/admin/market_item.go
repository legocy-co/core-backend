package admin

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	legoDomain "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/domain/marketplace/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
)

type MarketItemAdminCreateRequest struct {
	LegoSetID   int     `json:"legoSetID"`
	SellerID    int     `json:"sellerID"`
	Price       float32 `json:"price"`
	Location    string  `json:"location"`
	Status      string  `json:"status"`
	SetState    string  `json:"setState"`
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
	LegoSetID   int     `json:"legoSetID"`
	SellerID    int     `json:"sellerID"`
	Price       float32 `json:"price"`
	Location    string  `json:"location"`
	Status      string  `json:"status"`
	SetState    string  `json:"setState"`
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
	ID          int                         `json:"id"`
	Price       float32                     `json:"price"`
	Location    string                      `json:"location"`
	LegoSet     lego.LegoSetResponse        `json:"legoSet"`
	Seller      users.UserDetailResponse    `json:"seller"`
	Status      string                      `json:"status"`
	SetState    string                      `json:"setState"`
	Description string                      `json:"description"`
	Images      []marketplace.ImageResponse `json:"images"`
}

func GetMarketItemAdminResponse(
	mi *models.MarketItemAdmin) MarketItemAdminResponse {

	images := make([]marketplace.ImageResponse, 0, len(mi.Images))
	for _, img := range mi.Images {
		images = append(images, marketplace.GetImageResponse(img))
	}

	return MarketItemAdminResponse{
		ID:          mi.ID,
		Price:       mi.Price,
		Location:    mi.Location,
		LegoSet:     lego.GetLegoSetResponse(&mi.LegoSet),
		Seller:      users.GetUserDetailResponse(&mi.Seller),
		Status:      mi.Status,
		SetState:    mi.SetState,
		Description: mi.Description,
		Images:      images,
	}
}
