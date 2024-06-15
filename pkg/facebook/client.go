package facebook

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
)

func GetUserInfo(ctx context.Context, cfg *oauth2.Config, code string) (*TokenPayload, error) {
	token, err := cfg.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := cfg.Client(ctx, token)
	resp, err := client.Get("https://graph.facebook.com/me?fields=id,name,email")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	payload := &TokenPayload{}
	if err := json.NewDecoder(resp.Body).Decode(payload); err != nil {
		return nil, err
	}

	return payload, nil
}
