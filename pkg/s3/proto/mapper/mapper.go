package mapper

import (
	"github.com/docker/docker/image"
	"github.com/legocy-co/legocy/pkg/helpers"
	"github.com/legocy-co/legocy/pkg/s3/models"
	"github.com/legocy-co/legocy/pkg/s3/proto"
	"os"
)

func GetImageUploadRequest(file *os.File, bucketName string) *proto.UploadImageRequest {
	return &proto.UploadImageRequest{
		Meta: &proto.ImageInfo{
			Id:         int32(image.ID),
			BucketName: bucketName,
			FileFormat: file.,
		},
		Data: helpers.StreamToByte(file),
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
