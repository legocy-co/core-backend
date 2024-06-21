package fb

import (
	"crypto/sha256"
	"errors"
	"github.com/gin-gonic/gin"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/config"
)

var (
	cfg = config.GetAppConfig()
)

func checkSignUpKey(ctx *gin.Context, data schemas.FacebookSignUpRequest) error {

	key := ctx.GetHeader("X-Secret-Key")
	if key == "" {
		return errors.New("missing sign up key")
	}

	hashedValue := getHashedValue(data.Email, data.FacebookID)
	if key != hashedValue {
		return errors.New("invalid sign up key")
	}

	return nil
}

func checkSignInKey(ctx *gin.Context, data schemas.FacebookSignInRequest, user models.User) error {

	key := ctx.GetHeader("X-Secret-Key")
	if key == "" {
		return errors.New("missing sign up key")
	}

	hashedValue := getHashedValue(user.Email, data.FacebookID)
	if key != hashedValue {
		return errors.New("invalid sign up key")
	}

	return nil
}

func getHashedValue(email, id string) string {
	h := sha256.New()

	h.Write([]byte(email + id + cfg.FacebookSecretKeySalt))

	return string(h.Sum(nil))
}
