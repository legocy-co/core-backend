package marketplace

import (
	models "legocy-go/internal/domain/marketplace/models"
)

type LocationRequest struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func (loc *LocationRequest) ToLocationValueObject() *models.LocationValueObject {
	return &models.LocationValueObject{
		Country: loc.Country,
		City:    loc.City,
	}
}

type LocationResponse struct {
	ID      int    `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
}

func GetLocationResponse(loc *models.Location) LocationResponse {
	return LocationResponse{
		ID:      loc.ID,
		Country: loc.Country,
		City:    loc.City,
	}
}
