package marketplace

import (
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/domain/users/errors"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"strings"
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

type ImageDownloadRequest struct {
	ImagePath string `json:"imagePath"`
}

func (r ImageDownloadRequest) ToBucketNameImageName() (bucketName string, imageName string, err error) {
	fp := r.ImagePath
	if f := string(fp[0]); f == "/" {
		fp = fp[1:]
	}

	idx := strings.Index(fp, "/")
	if idx == len(fp)-1 {
		return "", "", errors.ErrInvalidImageFilepath
	}
	if idx < 0 || len(fp[idx+1:]) <= 0 {
		return "", "", errors.ErrInvalidImageFilepath
	}

	return fp[:idx], fp[idx+1:], nil
}

type ImageResponse struct {
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
	IsMain   bool   `json:"is_main"`
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
