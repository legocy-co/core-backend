package google

import (
	"context"
	"fmt"
	"github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/pkg/s3"
	types "github.com/legocy-co/legocy/pkg/s3/models"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"
)

func (h Handler) uploadImage(user *models.User, url string) error {

	if url == "" {
		return nil
	}

	file, err := getFile(url)
	if err != nil {
		return err
	}

	defer file.Close()
	slug, err := h.getSavedSlug(file, s3.UserBucketName, user.ID)
	if err != nil {
		return err
	}

	return h.imageService.StoreUserImage(
		context.Background(),
		&models.UserImage{UserID: user.ID, FilepathURL: slug},
	)

}

func getFile(url string) (*os.File, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	img, format, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	file, err := os.CreateTemp("", fmt.Sprintf("downloaded-%v", "."+format))
	if err != nil {
		return nil, err
	}

	switch strings.ToLower(format) {
	case "jpeg":
		err = jpeg.Encode(file, img, nil)
	case "png":
		err = png.Encode(file, img)
	}

	return file, err
}

func (h Handler) getSavedSlug(file *os.File, bucketName string, objectID int) (string, error) {

	fi, err := file.Stat()
	if err != nil {
		return "", err
	}

	// Image Type
	img := types.ImageUnitFromFile(file, objectID, file.Name(), fi.Size())

	// Save image to s3
	return h.imageStorage.UploadImage(img, bucketName)
}
