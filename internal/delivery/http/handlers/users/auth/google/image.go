package google

import (
	"context"
	"github.com/legocy-co/legocy/internal/pkg/s3"
	proto2 "github.com/legocy-co/legocy/internal/pkg/s3/proto"

	"github.com/legocy-co/legocy/internal/domain/users/models"
)

func (h Handler) uploadImage(user *models.User, url string) error {

	if url == "" {
		return nil
	}

	req := proto2.NewUploadImageURLRequest(
		proto2.UploadImageURLOpts{
			URL:      url,
			ObjectID: user.ID,
			Bucket:   s3.UserBucketName,
		},
	)

	slug, err := h.imageStorage.UploadImageFromURL(context.Background(), req)
	if err != nil {
		return err
	}

	return h.imageService.StoreUserImage(
		context.Background(), &models.UserImage{
			UserID:      user.ID,
			FilepathURL: slug,
		},
	)
}
