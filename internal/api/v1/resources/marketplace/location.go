package marketplace

import models "legocy-go/pkg/marketplace/models"

type LocationRequest struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func (loc *LocationRequest) ToLocationBasic() *models.LocationBasic {
	return &models.LocationBasic{
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
