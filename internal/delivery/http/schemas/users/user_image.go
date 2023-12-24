package users

import (
	"github.com/legocy-co/legocy/config"
	"github.com/legocy-co/legocy/internal/delivery/http/constants"
	"github.com/legocy-co/legocy/internal/domain/users/errors"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/pkg/helpers"
	"strings"
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

func NewUserDownloadImageRequest(query string) (*UserDownloadImageRequest, error) {
	if len(query) < len(constants.UserImagesBucketName) {
		return nil, errors.ErrInvalidImageFilepath
	}

	// Check if bucket
	if query[:5] != constants.UserImagesBucketName {
		return nil, errors.ErrInvalidImageFilepath
	}

	return &UserDownloadImageRequest{ImagePath: query}, nil
}

func (r UserDownloadImageRequest) ToBucketNameImageName() (bucketName string, imageName string, err error) {
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

type UserImagesListResponse struct {
	Images []UserImageInfoResponse `json:"images"`
}

func GetUserImagesListResponse(images []UserImageInfoResponse) UserImagesListResponse {
	return UserImagesListResponse{Images: images}
}

type UserImageInfoResponse struct {
	UserID      int    `json:"userID"`
	Filepath    string `json:"filepath"`
	DownloadURL string `json:"downloadURL"`
}

func GetUserInfoResponse(image *models.UserImage) UserImageInfoResponse {
	return UserImageInfoResponse{
		UserID:      image.UserID,
		Filepath:    image.FilepathURL,
		DownloadURL: config.GetAppConfig().BaseURL + "/api/v1/images/download?fp=" + helpers.EncodeURLString(image.FilepathURL),
	}
}
