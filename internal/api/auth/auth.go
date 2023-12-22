// internal/api/auth/auth.go
package auth

import (
	"net/http"
	"tech-test/config"

	"golang.org/x/oauth2"
)

// Config is the OAuth2 configuration
var Config = &oauth2.Config{
	ClientID:     config.GoogleClientID,
	ClientSecret: config.GoogleClientSecret,
	RedirectURL:  config.GoogleRedirectURL,
	Endpoint:     oauth2.GoogleEndpoint,
	Scopes:       []string{"profile", "email"},
}

// LoginHandler handles the login with Gmail
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	url := Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// CallbackHandler handles the OAuth callback
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Handle OAuth callback
	// ...

	// Redirect user to the desired page
	http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
}
