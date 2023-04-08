package users

import (
	"errors"
	"strings"
)

var (
	ErrInvalidImagePath = errors.New("invalid image path")
)

type UserImageUploadResponse struct {
	ImageURL string `json:"imageURL"`
	OK       bool   `json:"ok"`
}

func GetUserImageUploadResponse(imgUrl string) *UserImageUploadResponse {
	return &UserImageUploadResponse{
		ImageURL: imgUrl,
		OK:       imgUrl != "",
	}
}

type UserDownloadImageRequest struct {
	ImagePath string `json:"imagePath"`
}

func (r UserDownloadImageRequest) ToBucketNameImageName() (bucketName string, imageName string, err error) {
	idx := strings.Index(r.ImagePath, "/")
	if idx < 0 || len(r.ImagePath[idx+1:]) <= 0 {
		return "", "", ErrInvalidImagePath
	}

	return r.ImagePath[:idx], r.ImagePath[idx:], nil
}
