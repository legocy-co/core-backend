package marketplace

import (
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/domain/users/errors"
	"github.com/legocy-co/legocy/pkg/storage"
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
		ImageURL: imageURL,
	}
}

type ImageDownloadRequest struct {
	ImagePath string `json:"imagePath"`
}

func NewImageDownloadRequest(query string) (*ImageDownloadRequest, error) {
	if len(query) < len(storage.MarketItemsBucketName) {
		return nil, errors.ErrInvalidImageFilepath
	}

	// Check if bucket
	if query[:len(storage.MarketItemsBucketName)] != storage.MarketItemsBucketName {
		return nil, errors.ErrInvalidImageFilepath
	}

	return &ImageDownloadRequest{ImagePath: query}, nil
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