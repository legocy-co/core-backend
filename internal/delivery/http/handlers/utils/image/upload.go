package image

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/pkg/s3/client"
	image "github.com/legocy-co/legocy/pkg/s3/models"
	"strconv"
)

type UploadHandler func(ctx *gin.Context) (string, error)

func NewUploadHandler(s client.ImageStorage, bucketName, objectIdQueryParam string) UploadHandler {
	return func(ctx *gin.Context) (string, error) {
		// Get market item id
		objectId, err := strconv.Atoi(ctx.Param(objectIdQueryParam))
		if err != nil {
			return "", err
		}

		// Get file
		file, err := ctx.FormFile("file")
		if err != nil {
			return "", err
		}

		// Open file
		src, err := file.Open()
		defer src.Close()
		if err != nil {
			return "", err
		}

		// Domain Image
		img := image.ImageUnitFromFile(src, objectId, file.Filename, file.Size)

		// Save image to s3
		imgUrl, err := s.UploadImage(img, bucketName)
		if err != nil {
			return "", err
		}

		return imgUrl, nil
	}
}
