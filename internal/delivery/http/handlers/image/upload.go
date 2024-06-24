package image

import (
	"github.com/legocy-co/legocy/internal/pkg/s3/client"
	"github.com/legocy-co/legocy/internal/pkg/s3/proto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/lib/helpers"
)

type UploadHandler func(ctx *gin.Context) (string, error)

func NewUploadHandler(s client.ImageStorage, bucketName, objectIdQueryParam string) UploadHandler {
	return func(ctx *gin.Context) (string, error) {
		// Get object id
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

		// Save image to s3
		imgUrl, err := s.UploadImageFromFile(
			ctx, proto.NewUploadImageFileRequest(
				proto.UploadImageFileOpts{
					Data:     helpers.StreamToByte(src),
					ObjectID: objectId,
					Bucket:   bucketName,
					Format:   helpers.GetFileExtension(file.Filename),
				},
			),
		)

		return imgUrl, err
	}
}
