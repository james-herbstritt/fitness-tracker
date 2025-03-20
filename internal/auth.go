package internal

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func GenerateAuthCodeURL(verifier string) string {
	return FitbitOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
}

func ExchangeAuthCode(c *gin.Context, code string, verifier string) (*oauth2.Token, error) {
	return FitbitOAuthConfig.Exchange(c, code, oauth2.VerifierOption(verifier))
}

func GetTokenSource(c *gin.Context, token *oauth2.Token) oauth2.TokenSource {
	return FitbitOAuthConfig.TokenSource(c, token)
}

// TODO:: some error handling here
func GetToken(tokenSource oauth2.TokenSource) (*oauth2.Token, error) {
	return tokenSource.Token()
}
