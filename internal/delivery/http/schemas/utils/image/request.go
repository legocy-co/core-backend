package image

import (
	"github.com/legocy-co/legocy/internal/domain/users/errors"
	"github.com/legocy-co/legocy/internal/pkg/s3"
	"github.com/legocy-co/legocy/lib/helpers"
	"strings"
)

type DownloadImageRequest struct {
	imagePath string
}

func NewDownloadImageRequest(imagePath string) (*DownloadImageRequest, error) {

	bucketName := strings.Split(imagePath, "/")[0]

	if !helpers.Contains(s3.BucketNames, bucketName) {
		return nil, errors.ErrInvalidImageFilepath
	}

	return &DownloadImageRequest{imagePath: imagePath}, nil
}

func (r DownloadImageRequest) ToBucketNameImageName() (bucketName string, imageName string, err error) {
	fp := r.imagePath
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
