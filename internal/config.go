package internal

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/fitbit"
	"os"
)

var FitbitOAuthConfig *oauth2.Config = &oauth2.Config{
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	RedirectURL:  os.Getenv("FITBIT_REDIRECT_URI"),
	Scopes: []string{
		"activity",
		"heartrate",
		"nutrition",
		"oxygen_saturation",
		"respiratory_rate",
		"settings",
		"sleep",
		"temperature weight",
		"profile",
	},
	Endpoint: fitbit.Endpoint,
}
