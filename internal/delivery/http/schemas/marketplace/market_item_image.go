package marketplace

import (
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/config"
)

type ImageUploadResponse struct {
	OK       bool   `json:"ok"`
	ID       int    `json:"id"`
	ImageURL string `json:"imageURL"`
}

func NewImageUploadResponse(imageURL string, img *models.MarketItemImage) *ImageUploadResponse {
	return &ImageUploadResponse{
		OK:       imageURL != "",
		ID:       img.ID,
		ImageURL: config.GetAppConfig().BaseURL + "/api/v1/images/download?fp=" + imageURL,
	}
}

type ImageResponse struct {
	ID       int    `json:"id"`
	ImageURL string `json:"imageURL"`
	IsMain   bool   `json:"isMain"`
}

func GetImagesResponse(imgs []*models.MarketItemImage) []ImageResponse {
	images := make([]ImageResponse, 0, len(imgs))
	for _, img := range imgs {
		images = append(images, GetImageResponse(img))
	}
	return images
}

func GetImageResponse(img *models.MarketItemImage) ImageResponse {
	return ImageResponse{
		ID:       img.ID,
		ImageURL: config.GetAppConfig().CDNBaseURL + img.ImageURL,
		IsMain:   img.IsMain,
	}
}
