package google

import (
	"github.com/legocy-co/legocy/internal/domain/users/repository"
	"github.com/legocy-co/legocy/internal/domain/users/service"
	"github.com/legocy-co/legocy/internal/pkg/app"
	"github.com/legocy-co/legocy/internal/pkg/s3/client"
)

type Handler struct {
	r            repository.UserExternalAuthRepository
	imageService service.UserImageService
	imageStorage client.ImageStorage
}

func NewHandler(app *app.App) Handler {
	return Handler{
		r:            app.GetGoogleAuthRepository(),
		imageService: app.GetUserImagesService(),
		imageStorage: app.GetImageStorageClient(),
	}
}
