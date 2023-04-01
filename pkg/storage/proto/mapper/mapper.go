package mapper

import (
	"legocy-go/pkg/helpers"
	"legocy-go/pkg/storage/models"
	"legocy-go/pkg/storage/proto"
)

func GetImageRequest(image *models.ImageUnit, bucketName string) *proto.UploadImageRequest {
	return &proto.UploadImageRequest{
		Meta: &proto.ImageInfo{
			Id:         int32(image.ID),
			BucketName: bucketName,
		},
		Data: helpers.StreamToByte(image.Payload),
	}
}
