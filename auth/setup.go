package auth

import (
	"os"

	"github.com/stretchr/signature"

	"github.com/joho/godotenv"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
)

// Setup oauth providers for user authentication
func Setup() {
	godotenv.Load()
	gomniauth.SetSecurityKey(signature.RandomKey(16))
	gomniauth.WithProviders(
		facebook.New(
			os.Getenv("FACEBOOK_ID"),
			os.Getenv("FACEBOOK_SECRET"),
			os.Getenv("FACEBOOK_REDIRECT"),
		),
		github.New(
			os.Getenv("GITHUB_ID"),
			os.Getenv("GITHUB_SECRET"),
			os.Getenv("GITHUB_REDIRECT"),
		),
	)
}
