package google

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/idtoken"
	"os"
	"testing"
)

func TestTokenValidation(t *testing.T) {

	appID := os.Getenv("GOOGLE_CLIENT_ID")
	token := os.Getenv("GOOGLE_TEST_TOKEN")

	if token == "" {
		log.Println("empty token. Skipping test...")
		return
	}

	payload, err := idtoken.Validate(context.Background(), token, appID)
	if err != nil {
		t.Errorf(err.Error())
	}

	log.Printf("%+v", payload)
}
