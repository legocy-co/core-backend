package google

import (
	"context"
	"fmt"
	"google.golang.org/api/idtoken"
	"os"
	"testing"
)

func TestTokenValidation(t *testing.T) {

	appID := os.Getenv("GOOGLE_CLIENT_ID")
	token := os.Getenv("GOOGLE_TEST_TOKEN")

	if token == "" {
		t.Error("Token not found\n")
	}

	payload, err := idtoken.Validate(context.Background(), token, appID)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%+v", payload)
}
