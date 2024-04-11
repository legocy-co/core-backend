package lego

import (
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/pkg/config"
)

type LegoSetImageResponse struct {
	ID        int    `json:"id"`
	LegoSetID int    `json:"legoSetID"`
	IsMain    bool   `json:"isMain"`
	ImageURL  string `json:"imageURL"`
}

func GetLegoSetImageResponse(m *models.LegoSetImage) LegoSetImageResponse {
	return LegoSetImageResponse{
		ID:        m.ID,
		IsMain:    m.IsMain,
		LegoSetID: m.LegoSetID,
		ImageURL:  config.GetAppConfig().CDNBaseURL + m.ImageURL,
	}
}

func GetLegoSetImagesResponse(m []*models.LegoSetImage) []LegoSetImageResponse {
	var response []LegoSetImageResponse
	for _, v := range m {
		response = append(response, GetLegoSetImageResponse(v))
	}
	return response
}
