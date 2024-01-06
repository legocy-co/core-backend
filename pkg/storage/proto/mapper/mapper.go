package mapper

import (
	"github.com/legocy-co/legocy/pkg/helpers"
	"github.com/legocy-co/legocy/pkg/storage/models"
	"github.com/legocy-co/legocy/pkg/storage/proto"
)

func GetImageUploadRequest(image *models.ImageUnit, bucketName string) *proto.UploadImageRequest {
	return &proto.UploadImageRequest{
		Meta: &proto.ImageInfo{
			Id:         int32(image.ID),
			BucketName: bucketName,
			FileFormat: image.PayloadType,
		},
		Data: helpers.StreamToByte(image.Payload),
	}
}

func GetImageDownloadRequest(bucketName, imageName string) *proto.DownloadImageRequest {
	return &proto.DownloadImageRequest{
		BucketName: bucketName,
		ImageName:  imageName,
	}
}

func DownloadImageResponseToImageUnit(response *proto.DownloadImageResponse) *models.ImageUnit {
	// FIXME: Not enough metadata coming
	return &models.ImageUnit{
		ID:          0,
		Payload:     helpers.ByteToStream(response.Data),
		PayloadName: "",
		PayloadSize: int64(len(response.Data)),
		PayloadType: "image/png",
	}
}
