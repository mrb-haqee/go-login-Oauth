package model

import (
	"github.com/mrb-haqee/go-login-Oauth/helper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "189828178763-e2kpfrrgg8t92akp57c31p7eg64v9dpo.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	// Some random string, random for each request
	OauthStateString, _ = helper.GenerateRandomString()
)