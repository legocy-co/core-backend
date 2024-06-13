package facebook

import (
	"github.com/legocy-co/legocy/internal/pkg/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	cfg = config.GetAppConfig()
)

func GetOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.FacebookAppID,
		ClientSecret: cfg.FacebookSecret,
		RedirectURL:  cfg.FacebookCallbackURL,
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}
}

func GetSessionSecret() string {
	return cfg.FacebookSessionSecret
}
