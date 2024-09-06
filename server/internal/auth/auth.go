package auth

import (
	"log"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

const (
	Credentials = iota
	Github_OAuth
	Microsoft_OAuth
	Google_OAuth
)

// ? Should we move these to an env variable?
const (
	key    = "0bc89a6c-a7eb-4a6f-8f8f-425fe4ef9592"
	MaxAge = time.Hour * 24 * 30
	IsProd = false
)

var providersLookupMap = map[string]int{
	"google":      Google_OAuth,
	"github":      Github_OAuth,
	"microsoft":   Microsoft_OAuth,
	"credentials": Credentials,
}
var ProvidersMap = map[int]string{
	Google_OAuth:    "google",
	Github_OAuth:    "github",
	Microsoft_OAuth: "microsoft",
	Credentials:     "credentials",
}

func GetProviders() []string {
	providers := goth.GetProviders()
	var m []string

	for p, _ := range providers {
		m = append(m, p)
	}

	return m

}

func NewAuthConfig(provs []string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment variables")
	}

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(int(MaxAge))
	store.Options.Path = "/api/v1/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			os.Getenv("GOOGLE_CALLBACK_URL"),
		),
		github.New(
			os.Getenv("GITHUB_CLIENT_ID"),
			os.Getenv("GITHUB_CLIENT_SECRET"),
			os.Getenv("GITHUB_CALLBACK_URL"),
		),
	)

}
