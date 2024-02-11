package lego

import (
	"github.com/legocy-co/legocy/config"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSetImageResponse struct {
	ID        int    `json:"id"`
	LegoSetID int    `json:"lego_set_id"`
	IsMain    bool   `json:"is_main"`
	ImageURL  string `json:"image_url"`
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
