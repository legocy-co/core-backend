package google

import (
	"context"
	"log"
	"os"
	"testing"

	"google.golang.org/api/idtoken"
)

func TestTokenVerification(t *testing.T) {

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	token := os.Getenv("GOOGLE_TOKEN")

	payload, err := idtoken.Validate(context.Background(), token, clientID)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	log.Printf("%+v", payload)
}
